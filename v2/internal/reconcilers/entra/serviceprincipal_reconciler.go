// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package entra

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-logr/logr"
	msgraphsdk "github.com/microsoftgraph/msgraph-beta-sdk-go"
	msgraphmodels "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
	"github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals"
	"github.com/rotisserie/eris"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"

	asoentra "github.com/Azure/azure-service-operator/v2/api/entra/v1"
	"github.com/Azure/azure-service-operator/v2/internal/config"
	"github.com/Azure/azure-service-operator/v2/internal/reconcilers"
	"github.com/Azure/azure-service-operator/v2/internal/resolver"
	"github.com/Azure/azure-service-operator/v2/internal/util/kubeclient"
	"github.com/Azure/azure-service-operator/v2/pkg/common/annotations"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
)

// EntraServicePrincipalReconciler reconciles tenant service principals by application ID.
type EntraServicePrincipalReconciler struct {
	reconcilers.ReconcilerCommon
	ResourceResolver   *resolver.Resolver
	Config             config.Values
	EntraClientFactory EntraConnectionFactory
}

var _ genruntime.Reconciler = &EntraServicePrincipalReconciler{}

const servicePrincipalCreatedAnnotation = "serviceoperator.azure.com/service-principal-created"

func NewEntraServicePrincipalReconciler(
	kubeClient kubeclient.Client,
	entraClientFactory EntraConnectionFactory,
	resourceResolver *resolver.Resolver,
	positiveConditions *conditions.PositiveConditionBuilder,
	cfg config.Values,
) *EntraServicePrincipalReconciler {
	return &EntraServicePrincipalReconciler{
		ResourceResolver:   resourceResolver,
		Config:             cfg,
		EntraClientFactory: entraClientFactory,
		ReconcilerCommon: reconcilers.ReconcilerCommon{
			KubeClient:         kubeClient,
			PositiveConditions: positiveConditions,
		},
	}
}

func (r *EntraServicePrincipalReconciler) CreateOrUpdate(
	ctx context.Context,
	log logr.Logger,
	_ record.EventRecorder,
	obj genruntime.MetaObject,
	_ annotations.ResolvedReconcilePolicies,
) (ctrl.Result, error) {
	sp, err := r.asServicePrincipal(obj)
	if err != nil {
		return ctrl.Result{}, err
	}
	if (sp.Spec.AppId == nil || *sp.Spec.AppId == "") &&
		(sp.Spec.DisplayName == nil || *sp.Spec.DisplayName == "") {
		return ctrl.Result{}, eris.Errorf("service principal %s requires appId or displayName", sp.Name)
	}

	if id, ok := getEntraID(sp); ok {
		return r.update(ctx, id, sp, log)
	}

	if r.canAdopt(sp) {
		id, err := r.tryAdopt(ctx, sp, log)
		if err != nil {
			return ctrl.Result{}, eris.Wrapf(err, "trying to adopt service principal %s", sp.Name)
		}
		if id != "" {
			setEntraID(sp, id)
			genruntime.AddAnnotation(sp, servicePrincipalCreatedAnnotation, "false")
			return r.update(ctx, id, sp, log)
		}
	}

	if r.canCreate(sp) {
		if sp.Spec.AppId == nil || *sp.Spec.AppId == "" {
			return ctrl.Result{}, eris.Errorf("cannot create service principal %s without an appId; no principal with displayName %q was found", sp.Name, *sp.Spec.DisplayName)
		}
		return r.create(ctx, sp, log)
	}
	return ctrl.Result{}, eris.Errorf("service principal %s not found for adoption", sp.Name)
}

func (r *EntraServicePrincipalReconciler) Delete(
	ctx context.Context,
	_ logr.Logger,
	_ record.EventRecorder,
	obj genruntime.MetaObject,
) (ctrl.Result, error) {
	sp, err := r.asServicePrincipal(obj)
	if err != nil {
		return ctrl.Result{}, err
	}
	if !r.canCreate(sp) || sp.GetAnnotations()[servicePrincipalCreatedAnnotation] != "true" {
		return ctrl.Result{}, nil
	}
	id, ok := getEntraID(obj)
	if !ok {
		return ctrl.Result{}, nil
	}
	connection, err := r.EntraClientFactory(ctx, obj)
	if err != nil {
		return ctrl.Result{}, eris.Wrap(err, "creating entra client")
	}
	err = connection.Client().ServicePrincipals().ByServicePrincipalId(id).Delete(ctx, nil)
	if err != nil && !isNotFound(err) {
		return ctrl.Result{}, eris.Wrapf(err, "deleting service principal %s", id)
	}
	return ctrl.Result{}, nil
}

func (r *EntraServicePrincipalReconciler) Claim(
	_ context.Context,
	_ logr.Logger,
	_ record.EventRecorder,
	_ genruntime.MetaObject,
) error {
	return nil
}

func (r *EntraServicePrincipalReconciler) update(
	ctx context.Context,
	id string,
	sp *asoentra.ServicePrincipal,
	log logr.Logger,
) (ctrl.Result, error) {
	connection, err := r.EntraClientFactory(ctx, sp)
	if err != nil {
		return ctrl.Result{}, eris.Wrap(err, "creating entra client prior to update")
	}
	current, err := r.loadServicePrincipalByID(ctx, id, connection.Client())
	if err != nil {
		return ctrl.Result{}, eris.Wrapf(err, "getting service principal by ID %s", id)
	}
	if current == nil {
		setEntraID(sp, "")
		genruntime.AddAnnotation(sp, servicePrincipalCreatedAnnotation, "false")
		return ctrl.Result{Requeue: true}, nil
	}
	if sp.Spec.AppId != nil && (current.GetAppId() == nil || !strings.EqualFold(*current.GetAppId(), *sp.Spec.AppId)) {
		return ctrl.Result{}, eris.Errorf("service principal %s has appId %q, expected %q", id, valueOrEmpty(current.GetAppId()), *sp.Spec.AppId)
	}
	if sp.Spec.AppId == nil && (current.GetDisplayName() == nil || *current.GetDisplayName() != *sp.Spec.DisplayName) {
		return ctrl.Result{}, eris.Errorf("service principal %s has displayName %q, expected %q", id, valueOrEmpty(current.GetDisplayName()), *sp.Spec.DisplayName)
	}

	// Only principals created by ASO may be modified; appId is immutable in Graph.
	if sp.Spec.DisplayName != nil &&
		r.canCreate(sp) &&
		sp.GetAnnotations()[servicePrincipalCreatedAnnotation] == "true" {
		patch := msgraphmodels.NewServicePrincipal()
		patch.SetDisplayName(sp.Spec.DisplayName)
		updated, err := connection.Client().ServicePrincipals().ByServicePrincipalId(id).Patch(ctx, patch, nil)
		if err != nil {
			return ctrl.Result{}, eris.Wrapf(err, "updating service principal %s", id)
		}
		if updated == nil || updated.GetId() == nil || updated.GetAppId() == nil {
			updated, err = r.loadServicePrincipalByID(ctx, id, connection.Client())
			if err != nil {
				return ctrl.Result{}, eris.Wrapf(err, "reloading service principal %s after update", id)
			}
			if updated == nil {
				setEntraID(sp, "")
				genruntime.AddAnnotation(sp, servicePrincipalCreatedAnnotation, "false")
				return ctrl.Result{Requeue: true}, nil
			}
		}
		current = updated
	}
	sp.Status.AssignFromServicePrincipal(current)
	if err := r.saveAssociatedKubernetesResources(ctx, sp, log); err != nil {
		return ctrl.Result{}, eris.Wrapf(err, "saving associated Kubernetes resources for service principal %s", sp.Name)
	}
	return ctrl.Result{}, nil
}

func valueOrEmpty(value *string) string {
	if value == nil {
		return ""
	}
	return *value
}

func (r *EntraServicePrincipalReconciler) tryAdopt(
	ctx context.Context,
	sp *asoentra.ServicePrincipal,
	_ logr.Logger,
) (string, error) {
	connection, err := r.EntraClientFactory(ctx, sp)
	if err != nil {
		return "", eris.Wrap(err, "creating entra client prior to adoption search")
	}

	// The global application ID identifies the principal even if its display name changed.
	if sp.Spec.AppId != nil && *sp.Spec.AppId != "" {
		matches, err := r.findServicePrincipals(ctx, "appId", *sp.Spec.AppId, connection.Client())
		if err != nil {
			return "", err
		}
		if len(matches) > 1 {
			return "", eris.Errorf("multiple service principals found with appId %s", *sp.Spec.AppId)
		}
		if len(matches) == 1 {
			return servicePrincipalID(matches[0], "appId", *sp.Spec.AppId)
		}
	}

	if sp.Spec.DisplayName != nil && *sp.Spec.DisplayName != "" {
		matches, err := r.findServicePrincipals(ctx, "displayName", *sp.Spec.DisplayName, connection.Client())
		if err != nil {
			return "", err
		}
		if len(matches) > 1 {
			return "", eris.Errorf("cannot adopt as multiple existing Entra service principals found with display name %q", *sp.Spec.DisplayName)
		}
		if len(matches) == 1 {
			principal := matches[0]
			if sp.Spec.AppId != nil && (principal.GetAppId() == nil || !strings.EqualFold(*principal.GetAppId(), *sp.Spec.AppId)) {
				return "", eris.Errorf("service principal with display name %q has appId %q, expected %q", *sp.Spec.DisplayName, valueOrEmpty(principal.GetAppId()), *sp.Spec.AppId)
			}
			return servicePrincipalID(principal, "display name", *sp.Spec.DisplayName)
		}
	}

	return "", nil
}

func servicePrincipalID(principal msgraphmodels.ServicePrincipalable, field, value string) (string, error) {
	if principal == nil {
		return "", eris.Errorf("service principal with %s %q was nil", field, value)
	}
	id := principal.GetId()
	if id == nil || *id == "" {
		return "", eris.Errorf("service principal with %s %q has no object ID", field, value)
	}
	return *id, nil
}

func (r *EntraServicePrincipalReconciler) findServicePrincipals(
	ctx context.Context,
	field string,
	value string,
	graph *msgraphsdk.GraphServiceClient,
) ([]msgraphmodels.ServicePrincipalable, error) {
	filter := fmt.Sprintf("%s eq '%s'", field, escapeODataString(value))
	options := &serviceprincipals.ServicePrincipalsRequestBuilderGetRequestConfiguration{
		QueryParameters: &serviceprincipals.ServicePrincipalsRequestBuilderGetQueryParameters{Filter: &filter},
	}
	builder := graph.ServicePrincipals()
	var matches []msgraphmodels.ServicePrincipalable
	for {
		result, err := builder.Get(ctx, options)
		if err != nil {
			return nil, eris.Wrapf(err, "searching service principals by %s %q", field, value)
		}
		if result == nil {
			return nil, eris.Errorf("searching service principals by %s %q returned no response", field, value)
		}
		matches = append(matches, result.GetValue()...)
		if len(matches) > 1 {
			return matches, nil
		}
		next := result.GetOdataNextLink()
		if next == nil || *next == "" {
			return matches, nil
		}
		builder = graph.ServicePrincipals().WithUrl(*next)
		options = nil
	}
}

func (r *EntraServicePrincipalReconciler) create(
	ctx context.Context,
	sp *asoentra.ServicePrincipal,
	log logr.Logger,
) (ctrl.Result, error) {
	connection, err := r.EntraClientFactory(ctx, sp)
	if err != nil {
		return ctrl.Result{}, eris.Wrap(err, "creating entra client prior to creation")
	}
	body := msgraphmodels.NewServicePrincipal()
	sp.Spec.AssignToServicePrincipal(body)
	result, err := connection.Client().ServicePrincipals().Post(ctx, body, nil)
	if err != nil {
		return ctrl.Result{}, eris.Wrapf(err, "creating service principal %s", sp.Name)
	}
	if result == nil || result.GetId() == nil || *result.GetId() == "" {
		return ctrl.Result{}, eris.Errorf("creating service principal %s: no object ID returned", sp.Name)
	}
	sp.Status.AssignFromServicePrincipal(result)
	setEntraID(sp, *result.GetId())
	genruntime.AddAnnotation(sp, servicePrincipalCreatedAnnotation, "true")
	if err := r.saveAssociatedKubernetesResources(ctx, sp, log); err != nil {
		return ctrl.Result{}, eris.Wrapf(err, "saving associated Kubernetes resources for service principal %s", sp.Name)
	}
	return ctrl.Result{}, nil
}

func (r *EntraServicePrincipalReconciler) UpdateStatus(
	ctx context.Context,
	log logr.Logger,
	_ record.EventRecorder,
	obj genruntime.MetaObject,
	_ annotations.ResolvedReconcilePolicies,
) error {
	sp, err := r.asServicePrincipal(obj)
	if err != nil {
		return err
	}
	id, ok := getEntraID(sp)
	if !ok {
		return nil
	}
	connection, err := r.EntraClientFactory(ctx, sp)
	if err != nil {
		return eris.Wrap(err, "creating entra client")
	}
	current, err := r.loadServicePrincipalByID(ctx, id, connection.Client())
	if err != nil {
		return eris.Wrapf(err, "getting service principal by ID %s", id)
	}
	if current == nil {
		return nil
	}
	sp.Status.AssignFromServicePrincipal(current)
	return r.saveAssociatedKubernetesResources(ctx, sp, log)
}

func (r *EntraServicePrincipalReconciler) loadServicePrincipalByID(
	ctx context.Context,
	id string,
	graph *msgraphsdk.GraphServiceClient,
) (msgraphmodels.ServicePrincipalable, error) {
	result, err := graph.ServicePrincipals().ByServicePrincipalId(id).Get(ctx, nil)
	if isNotFound(err) {
		return nil, nil
	}
	return result, err
}

func (r *EntraServicePrincipalReconciler) asServicePrincipal(
	obj genruntime.MetaObject,
) (*asoentra.ServicePrincipal, error) {
	sp, ok := obj.(*asoentra.ServicePrincipal)
	if !ok {
		return nil, eris.Errorf("cannot modify resource that is not of type *entra.ServicePrincipal. Type is %T", obj)
	}
	return sp, nil
}

func (r *EntraServicePrincipalReconciler) canAdopt(sp *asoentra.ServicePrincipal) bool {
	return sp.Spec.OperatorSpec == nil || sp.Spec.OperatorSpec.AdoptionAllowed()
}

func (r *EntraServicePrincipalReconciler) canCreate(sp *asoentra.ServicePrincipal) bool {
	return sp.Spec.OperatorSpec == nil || sp.Spec.OperatorSpec.CreationAllowed()
}

func (r *EntraServicePrincipalReconciler) saveAssociatedKubernetesResources(
	ctx context.Context,
	sp *asoentra.ServicePrincipal,
	_ logr.Logger,
) error {
	if sp.Spec.OperatorSpec == nil || sp.Spec.OperatorSpec.ConfigMaps == nil {
		return nil
	}
	destinations := sp.Spec.OperatorSpec.ConfigMaps
	collector := configmaps.NewCollector(sp.Namespace)
	if destinations.EntraID != nil && sp.Status.EntraID != nil {
		collector.AddValue(destinations.EntraID, *sp.Status.EntraID)
	}
	if destinations.AppId != nil && sp.Status.AppId != nil {
		collector.AddValue(destinations.AppId, *sp.Status.AppId)
	}
	values, err := collector.Values()
	if err != nil {
		return eris.Wrap(err, "collecting configmaps for service principal")
	}
	if len(values) == 0 {
		return nil
	}
	resources := configmaps.SliceToClientObjectSlice(values)
	_, err = genruntime.ApplyObjsAndEnsureOwner(ctx, r.KubeClient, sp, resources)
	return err
}
