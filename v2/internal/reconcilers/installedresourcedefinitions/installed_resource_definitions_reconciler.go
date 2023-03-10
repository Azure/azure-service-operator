/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package installedresourcedefinitions

import (
	"context"
	"os"
	"sync"

	"github.com/go-logr/logr"
	_ "github.com/go-sql-driver/mysql" //mysql driver
	"github.com/pkg/errors"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	serviceoperator "github.com/Azure/azure-service-operator/v2/api/serviceoperator/v1api"
	"github.com/Azure/azure-service-operator/v2/internal/config"
	"github.com/Azure/azure-service-operator/v2/internal/crdmanagement"
	. "github.com/Azure/azure-service-operator/v2/internal/logging"
	"github.com/Azure/azure-service-operator/v2/internal/reconcilers"
	"github.com/Azure/azure-service-operator/v2/internal/util/kubeclient"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
)

var _ genruntime.Reconciler = &InstalledResourceReconciler{}

type InstalledResourceReconciler struct {
	reconcilers.ReconcilerCommon
	Config config.Values

	mut  sync.Mutex
	crds []apiextensions.CustomResourceDefinition
}

func NewInstalledResourceDefinitionsReconciler(
	kubeClient kubeclient.Client,
	positiveConditions *conditions.PositiveConditionBuilder,
	cfg config.Values) *InstalledResourceReconciler {

	return &InstalledResourceReconciler{
		Config: cfg,
		ReconcilerCommon: reconcilers.ReconcilerCommon{
			KubeClient:         kubeClient,
			PositiveConditions: positiveConditions,
		},
	}
}

func (r *InstalledResourceReconciler) asInstalledResources(obj genruntime.MetaObject) (*serviceoperator.InstalledResourceDefinitions, error) {
	typedObj, ok := obj.(*serviceoperator.InstalledResourceDefinitions)
	if !ok {
		return nil, errors.Errorf("cannot modify resource that is not of type *serviceoperator.InstalledResourceDefinitions. Type is %T", obj)
	}

	return typedObj, nil
}

func (r *InstalledResourceReconciler) Claim(_ context.Context, _ logr.Logger, _ record.EventRecorder, _ genruntime.MetaObject) error {
	// We don't need to do anything to claim this resource
	return nil
}

func (r *InstalledResourceReconciler) CreateOrUpdate(ctx context.Context, log logr.Logger, eventRecorder record.EventRecorder, obj genruntime.MetaObject) (ctrl.Result, error) {
	_, err := r.asInstalledResources(obj)
	if err != nil {
		return ctrl.Result{}, err
	}

	crdManager := crdmanagement.NewManager(log, r.KubeClient)
	goalCRDs, err := r.loadCRDs(crdManager)
	if err != nil {
		return ctrl.Result{}, err
	}
	log.V(Info).Info("Goal CRDs", "count", len(goalCRDs))

	existingCRDs, err := crdManager.ListOperatorCRDs(ctx)
	if err != nil {
		return ctrl.Result{}, err
	}
	log.V(Info).Info("Existing CRDs", "count", len(existingCRDs))

	goalCRDsWithDifferentVersion := crdManager.FindGoalCRDsNeedingUpdate(existingCRDs, goalCRDs, crdmanagement.VersionEqual)
	goalCRDsWithDifferentSpec := crdManager.FindGoalCRDsNeedingUpdate(existingCRDs, goalCRDs, crdmanagement.SpecEqual)

	// The same CRD may be in both sets, but we don't want to apply twice, so combine the sets
	crdsToApply := make(map[string]apiextensions.CustomResourceDefinition)
	for name, crd := range goalCRDsWithDifferentVersion {
		crdsToApply[name] = crd
	}
	for name, crd := range goalCRDsWithDifferentSpec {
		crdsToApply[name] = crd
	}

	if len(crdsToApply) > 0 {
		for _, crd := range crdsToApply {
			crd := crd
			toApply := &apiextensions.CustomResourceDefinition{
				ObjectMeta: metav1.ObjectMeta{
					Name: crd.Name,
				},
			}
			result, err := controllerutil.CreateOrUpdate(ctx, r.KubeClient, toApply, func() error {
				resourceVersion := toApply.ResourceVersion
				*toApply = crd
				toApply.ResourceVersion = resourceVersion

				log.V(Verbose).Info("Applying CRD", "name", toApply.Name) // TODO: Delete this
				return nil
			})
			if err != nil {
				return ctrl.Result{}, errors.Wrapf(err, "failed to apply CRD %s", crd.Name)
			}

			log.V(Debug).Info("Successfully applied CRD", "name", crd.Name, "result", result)
		}

		// If we make it to here, we have successfully updated all the CRDs we needed to. We need to kill the pod and let it restart so
		// that the new shape CRDs can be reconciled.
		log.V(Status).Info("Restarting operator pod after updating CRDs", "count", len(crdsToApply))
		os.Exit(0) // TODO: Should this be nonzero?
	}

	log.V(Status).Info("Successfully reconciled InstalledResourceDefinitions CRDs.", "count", len(existingCRDs))
	return ctrl.Result{}, nil
}

func (r *InstalledResourceReconciler) loadCRDs(crdManager *crdmanagement.Manager) ([]apiextensions.CustomResourceDefinition, error) {
	if len(r.crds) > 0 {
		// Nothing to do as they're already loaded. Pod has to restart for them to change
		return r.crds, nil
	}

	r.mut.Lock()
	defer r.mut.Unlock()

	crds, err := crdManager.LoadOperatorCRDs(crdmanagement.CRDLocation)
	if err != nil {
		return nil, err
	}
	crds = crdManager.FixCRDNamespaceRefs(crds, r.Config.PodNamespace)

	r.crds = crds
	return crds, nil
}

func (r *InstalledResourceReconciler) Delete(_ context.Context, _ logr.Logger, _ record.EventRecorder, _ genruntime.MetaObject) (ctrl.Result, error) {
	// We don't uninstall resources currently, so there's nothing to do here
	return ctrl.Result{}, nil
}

func (r *InstalledResourceReconciler) UpdateStatus(ctx context.Context, log logr.Logger, eventRecorder record.EventRecorder, obj genruntime.MetaObject) error {
	// TODO: Do we need anything on status?

	return nil
}
