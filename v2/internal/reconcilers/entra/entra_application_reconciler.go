/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package entra

import (
	"context"
	"fmt"
	"net/http"

	. "github.com/Azure/azure-service-operator/v2/internal/logging"

	"github.com/go-logr/logr"
	msgraphsdkgo "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/microsoftgraph/msgraph-sdk-go/applications"
	msgraphmodels "github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
	"github.com/rotisserie/eris"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	asoentra "github.com/Azure/azure-service-operator/v2/api/entra/v1"
	"github.com/Azure/azure-service-operator/v2/internal/config"
	"github.com/Azure/azure-service-operator/v2/internal/identity"
	"github.com/Azure/azure-service-operator/v2/internal/reconcilers"
	"github.com/Azure/azure-service-operator/v2/internal/resolver"
	"github.com/Azure/azure-service-operator/v2/internal/util/kubeclient"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
)

// EntraApplicationReconciler reconciles an Entra application.
type EntraApplicationReconciler struct {
	reconcilers.ReconcilerCommon
	ResourceResolver   *resolver.Resolver
	CredentialProvider identity.CredentialProvider
	Config             config.Values
	EntraClientFactory EntraConnectionFactory
}

var _ genruntime.Reconciler = &EntraApplicationReconciler{}

func NewEntraApplicationReconciler(
	kubeClient kubeclient.Client,
	entraClientFactory EntraConnectionFactory,
	resourceResolver *resolver.Resolver,
	positiveConditions *conditions.PositiveConditionBuilder,
	cfg config.Values,
) *EntraApplicationReconciler {
	return &EntraApplicationReconciler{
		ResourceResolver:   resourceResolver,
		Config:             cfg,
		EntraClientFactory: entraClientFactory,
		ReconcilerCommon: reconcilers.ReconcilerCommon{
			KubeClient:         kubeClient,
			PositiveConditions: positiveConditions,
		},
	}
}

func (r *EntraApplicationReconciler) CreateOrUpdate(
	ctx context.Context,
	log logr.Logger,
	eventRecorder record.EventRecorder,
	obj genruntime.MetaObject,
) (ctrl.Result, error) {
	app, err := r.asApplication(obj)
	if err != nil {
		return ctrl.Result{}, eris.Wrapf(err, "creating or updating application %s", app.Name)
	}

	// If we already know the Entra ID of the application (captured in an annotation), we can update it directly
	if id, ok := getEntraID(obj); ok {
		return r.update(ctx, id, app, log)
	}

	// If we're allowed to adopt the application, we can try to find it
	if r.canAdopt(app) {
		id, err := r.tryAdopt(ctx, app, log)
		if err != nil {
			return ctrl.Result{}, eris.Wrapf(err, "trying to adopt application %s", app.Name)
		}

		if id != "" {
			// We found an existing application to adopt
			setEntraID(obj, id)
			return r.update(ctx, id, app, log)
		}
	}

	// If we can't adopt, or didn't find an application to adopt, we can create a new one
	if r.canCreate(app) {
		return r.create(ctx, app, log)
	}

	// Nothing to do
	return ctrl.Result{}, nil
}

func (r *EntraApplicationReconciler) Delete(
	ctx context.Context,
	log logr.Logger,
	eventRecorder record.EventRecorder,
	obj genruntime.MetaObject,
) (ctrl.Result, error) {
	log.V(Status).Info("Deleting Entra application")

	app, err := r.asApplication(obj)
	if err != nil {
		return ctrl.Result{}, eris.Wrapf(err, "deleting application %s", app.Name)
	}

	// If we don't know the Entra ID of the application (captured in an annotation), there's nothing to do.
	id, ok := getEntraID(obj)
	if !ok {
		return ctrl.Result{}, nil
	}

	client, err := r.EntraClientFactory(ctx, obj)
	if err != nil {
		return ctrl.Result{}, eris.Wrap(err, "creating entra client")
	}

	err = client.Client().Applications().ByApplicationId(id).Delete(ctx, nil)
	if err != nil {
		// If the application doesn't exist, return nil as we've successfully ensured that it doesn't exist
		if r.isNotFound(err) {
			return ctrl.Result{}, nil
		}

		var displayName string
		if app.Spec.DisplayName != nil {
			displayName = *app.Spec.DisplayName
		}

		return ctrl.Result{}, eris.Wrapf(err, "failed to delete application %q (%s)", displayName, id)
	}

	return ctrl.Result{}, nil
}

func (r *EntraApplicationReconciler) Claim(
	ctx context.Context,
	log logr.Logger,
	eventRecorder record.EventRecorder,
	obj genruntime.MetaObject,
) error {
	// Nothing to do
	return nil
}

// update completes our reconciliation by updating an existing Entra application.
func (r *EntraApplicationReconciler) update(
	ctx context.Context,
	id string,
	app *asoentra.Application,
	log logr.Logger,
) (ctrl.Result, error) {
	log = log.WithValues("id", id)
	log.V(Status).Info("Updating Entra application")

	// Create our Entra Client
	client, err := r.EntraClientFactory(ctx, app)
	if err != nil {
		return ctrl.Result{}, eris.Wrap(err, "creating entra client prior to update")
	}

	// Load the existing application by ID
	a, err := r.loadApplicationByID(ctx, id, client.Client())
	if err != nil {
		if r.isNotFound(err) {
			// Application used to exist, but no longer does - it's probably been deleted
			// Remove the existing annotation and requeue the reconciliation to create a replacement
			log.V(Status).Info("Application no longer exists")
			setEntraID(app, "")
			return ctrl.Result{
				Requeue: true,
			}, nil
		}

		return ctrl.Result{}, eris.Wrapf(err, "getting application by ID %s", id)
	}

	// Update - PATCH
	app.Spec.AssignToApplication(a)

	_, err = client.Client().Applications().ByApplicationId(id).Patch(ctx, a, nil)
	if err != nil {
		// Failed to update
		return ctrl.Result{}, eris.Wrapf(err, "failed to update application %s", id)
	}

	app.Status.AssignFromApplication(a)

	return ctrl.Result{}, nil
}

// tryAdopt tries to find an existing Entra application to adopt.
// Returns the Entra ID of the application to adopt, and nil if found;
// nil, nil if no application was found to adopt;
// or nil, error if there was a problem finding the application.
func (r *EntraApplicationReconciler) tryAdopt(
	ctx context.Context,
	app *asoentra.Application,
	log logr.Logger,
) (string, error) {
	log.V(Status).Info("Searching for existing Entra application to adopt", "application", app.Name)

	// Create our Entra Client
	client, err := r.EntraClientFactory(ctx, app)
	if err != nil {
		return "", eris.Wrap(err, "creating entra client prior to adoption search")
	}

	// Try to find the application by DisplayName
	displayName := app.Spec.DisplayName
	if to.Value(displayName) == "" {
		// Can't adopt without a display name
		return "", nil
	}

	log.V(Status).Info("Searching for existing Entra application by display name", "displayName", *displayName)
	apps, err := r.loadApplicationsByDisplayName(ctx, *displayName, client.Client())
	if err != nil {
		if r.isNotFound(err) {
			// No application to adopt
			return "", nil
		}

		return "", eris.Wrapf(err, "getting application by display name %s", *displayName)
	}

	if len(apps) == 0 {
		// No application to adopt
		log.V(Status).Info("No existing Entra application found by display name", "displayName", *displayName)
		return "", nil
	}

	if len(apps) > 1 {
		// Multiple applications found with the same display name
		log.V(Status).Info("Multiple existing Entra applications found by display name", "displayName", *displayName)
		return "", eris.Errorf("multiple existing Entra applications found with display name %s", *displayName)
	}

	// We found a single application to adopt
	a := apps[0]
	id := a.GetId()
	if id == nil {
		return "", nil
	}

	return *id, nil
}

// create completes our reconciliation by creating a new Entra application.
func (r *EntraApplicationReconciler) create(
	ctx context.Context,
	app *asoentra.Application,
	log logr.Logger,
) (ctrl.Result, error) {
	log.V(Status).Info("Creating Entra application")

	// Create our Entra Client
	client, err := r.EntraClientFactory(ctx, app)
	if err != nil {
		return reconcile.Result{}, eris.Wrap(err, "creating entra client prior to creation")
	}

	a := msgraphmodels.NewApplication()
	app.Spec.AssignToApplication(a)

	status, err := client.Client().Applications().Post(ctx, a, nil)
	if err != nil {
		// Failed to create
		return ctrl.Result{}, eris.Wrapf(err, "failed to create application %s", app.Name)
	}

	app.Status.AssignFromApplication(status)

	if id := status.GetId(); id != nil {
		setEntraID(app, *id)
	}

	err = r.saveAssociatedKubernetesResources(ctx, app, log)
	if err != nil {
		return ctrl.Result{}, eris.Wrapf(err, "failed to save associated Kubernetes resources for application %s", app.Name)
	}

	return ctrl.Result{}, nil
}

func (r *EntraApplicationReconciler) UpdateStatus(
	ctx context.Context,
	log logr.Logger,
	eventRecorder record.EventRecorder,
	obj genruntime.MetaObject,
) error {
	app, err := r.asApplication(obj)
	if err != nil {
		return eris.Wrapf(err, "updating status of application %s", app.Name)
	}

	client, err := r.EntraClientFactory(ctx, obj)
	if err != nil {
		return eris.Wrap(err, "creating entra client")
	}

	id, ok := getEntraID(obj)
	if !ok {
		// If we don't know the Entra ID of the application, there's nothing to do.
		log.V(Status).Info("No Entra ID found for application, skipping status update")
		return nil
	}

	applicationable, err := r.loadApplicationByID(ctx, id, client.Client())
	if err != nil {
		// If the application doesn't exist, nothing to do as we're probably in the midst of deleting it
		if r.isNotFound(err) {
			return nil
		}

		// If the error is not a 404, return the error
		return eris.Wrapf(err, "failed to update status of application %s", id)
	}

	app.Status.AssignFromApplication(applicationable)

	err = r.saveAssociatedKubernetesResources(ctx, app, log)
	if err != nil {
		return eris.Wrapf(err, "failed to save associated Kubernetes resources for application %s", app.Name)
	}

	return nil
}

// loadApplicationByID loads the application from Entra, if it exists.
// Returns the application and nil if it exists, nil and nil if it doesn't, or nil and an error if there was a problem.
func (r *EntraApplicationReconciler) loadApplicationByID(
	ctx context.Context,
	id string,
	client *msgraphsdkgo.GraphServiceClient,
) (msgraphmodels.Applicationable, error) {
	// Try to get the application by ID
	applicationable, err := client.Applications().ByApplicationId(id).Get(ctx, nil)
	if err != nil {
		// If the only problem is that the application doesn't exist, return nil and nil
		if r.isNotFound(err) {
			return nil, nil
		}

		return nil, err
	}

	return applicationable, nil
}

// loadApplicationsByDisplayName loads applications from Entra by display name.
func (r *EntraApplicationReconciler) loadApplicationsByDisplayName(
	ctx context.Context,
	displayName string,
	client *msgraphsdkgo.GraphServiceClient,
) ([]msgraphmodels.Applicationable, error) {
	// Try to get the application by display name
	filterStr := fmt.Sprintf("displayName eq '%s'", displayName)

	query := &applications.ApplicationsRequestBuilderGetQueryParameters{
		Filter: &filterStr,
	}

	options := &applications.ApplicationsRequestBuilderGetRequestConfiguration{
		QueryParameters: query,
	}

	result, err := client.Applications().Get(ctx, options)
	if err != nil {
		return nil, eris.Wrapf(err, "failed to search for application with name %s", displayName)
	}

	apps := result.GetValue()
	return apps, nil
}

// isNotFound returns true if the error is a 404 error.
func (r *EntraApplicationReconciler) isNotFound(err error) bool {
	var odataError *odataerrors.ODataError
	if eris.As(err, &odataError) {
		if odataError.ResponseStatusCode == http.StatusNotFound {
			return true
		}
	}

	return false
}

func (r *EntraApplicationReconciler) asApplication(
	obj genruntime.MetaObject,
) (*asoentra.Application, error) {
	typedObj, ok := obj.(*asoentra.Application)
	if !ok {
		return nil, eris.Errorf("cannot modify resource that is not of type *entra.Application. Type is %T", obj)
	}

	return typedObj, nil
}

func (r *EntraApplicationReconciler) canAdopt(app *asoentra.Application) bool {
	if app.Spec.OperatorSpec == nil {
		// Default is AdoptOrCreate
		return true
	}

	return app.Spec.OperatorSpec.AdoptionAllowed()
}

func (r *EntraApplicationReconciler) canCreate(app *asoentra.Application) bool {
	if app.Spec.OperatorSpec == nil {
		// Default is AdoptOrCreate
		return true
	}

	return app.Spec.OperatorSpec.CreationAllowed()
}

// saveAssociatedKubernetesResources retrieves Kubernetes resources to create and saves them to Kubernetes.
// If there are no resources to save this method is a no-op.
func (r *EntraApplicationReconciler) saveAssociatedKubernetesResources(
	ctx context.Context,
	app *asoentra.Application,
	log logr.Logger,
) error {
	if app == nil ||
		app.Spec.OperatorSpec == nil {
		// No OperatorSpec, nothing to do
		return nil
	}

	// Accumulate all the resources we need to save
	var resources []client.Object

	operatorSpec := app.Spec.OperatorSpec
	if operatorSpec.ConfigMaps != nil {
		// If we have configmaps to export, we need to collect them
		collector := configmaps.NewCollector(app.Namespace)

		if operatorSpec.ConfigMaps.EntraID != nil && app.Status.EntraID != nil {
			// If we have an Entra ID configmap, we need to collect it
			collector.AddValue(operatorSpec.ConfigMaps.EntraID, *app.Status.EntraID)
		}

		if operatorSpec.ConfigMaps.AppId != nil && app.Status.AppId != nil {
			// If we have an AppId configmap, we need to collect it
			collector.AddValue(operatorSpec.ConfigMaps.AppId, *app.Status.AppId)
		}

		values, err := collector.Values()
		if err != nil {
			return eris.Wrap(err, "failed to collect configmaps for Entra application")
		}

		if len(values) > 0 {
			resources = append(resources, configmaps.SliceToClientObjectSlice(values)...)
		}
	}

	if len(resources) == 0 {
		// No resources to save, nothing to do
		return nil
	}

	// Save the resources to Kubernetes
	results, err := genruntime.ApplyObjsAndEnsureOwner(ctx, r.KubeClient, app, resources)
	if err != nil {
		return err
	}

	for i := 0; i < len(resources); i++ {
		resource := resources[i]
		result := results[i]

		log.V(Debug).Info("Successfully created resource",
			"namespace", resource.GetNamespace(),
			"name", resource.GetName(),
			"type", fmt.Sprintf("%T", resource),
			"action", result)
	}

	return nil
}
