package cont

import (
	"fmt"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"

	"context"

	"github.com/go-logr/logr"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

type ProvisionState string

type ResourceClient interface {
	Create(context.Context, runtime.Object) error
	Validate(context.Context, runtime.Object) (bool, error)
	Delete(context.Context, runtime.Object) error
}

type Definition struct {
	ProvisionState ProvisionState
	Name   string
	Spec   runtime.Object
	IsBeingDeleted bool
	AddFinalizer func(string)
	RemoveFinalizer func(string)
	HasFinalizer func(string) bool
}

type DefinitionFetcher interface {
	GetDefinition(ctx context.Context, kubeClient client.Client, req ctrl.Request) (Definition, error)
	GetDependencies(ctx context.Context, kubeClient client.Client, req ctrl.Request) ([]Definition, error)
}

// AzureController reconciles a ResourceGroup object
type AzureController struct {
	KubeClient     client.Client
	Log            logr.Logger
	Recorder       record.EventRecorder
	ResourceClient ResourceClient
	DefinitionFetcher DefinitionFetcher
	FinalizerName  string
}

type AzureControllerFactory interface {
	Create(kubeClient client.Client, log logr.Logger, recorder record.EventRecorder) *AzureController
}

// +kubebuilder:rbac:groups=azure.microsoft.com,resources=resourcegroups,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources=resourcegroups/status,verbs=get;update;patch

// Reconcile function does the main reconciliation loop of the operator
func (r *AzureController) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("resourcegroup", req.NamespacedName)

	definition, err := r.DefinitionFetcher.GetDefinition(ctx, r.KubeClient, req)
	if err != nil {
		log.Info("Unable to retrieve resourcegroup resource", "err", err.Error())
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	if definition.IsBeingDeleted {
		err := r.handleFinalizer(definition, r.FinalizerName)
		if err != nil {
			return reconcile.Result{}, fmt.Errorf("error when handling finalizer: %v", err)
		}
		return ctrl.Result{}, nil
	}

	if !definition.HasFinalizer(r.FinalizerName) {
		err := r.addFinalizer(definition, r.FinalizerName)
		if err != nil {
			return ctrl.Result{}, fmt.Errorf("error when removing finalizer: %v", err)
		}
		return ctrl.Result{}, nil
	}

	if definition.ProvisionState != "Provisioning" { // TODO
		//if !instance.IsSubmitted() {
		err := r.reconcileExternal(definition)
		if err != nil {
			return ctrl.Result{}, fmt.Errorf("error when creating resource in azure: %v", err)
		}
		return ctrl.Result{}, nil
	}

	return ctrl.Result{}, nil

}

// SetupWithManager function sets up the functions with the controller
func (r *AzureController) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&azurev1alpha1.ResourceGroup{}).
		Complete(r)
}

func (r *AzureController) reconcileExternal(definition Definition) error {

	ctx := context.Background()
	var err error
	resourcegroupName := instance.ObjectMeta.Name

	// write information back to instance
	instance.Status.Provisioning = true
	instance.Status.ProvisionState = azurev1alpha1.Provisioning
	err = r.KubeClient.Update(ctx, instance)
	if err != nil {
		//log error and kill it
		r.Recorder.Event(instance, "Warning", "Failed", "Unable to update instance")
	}

	err = r.ResourceClient.Create(ctx, instance)
	if err != nil {

		r.Recorder.Event(instance, "Warning", "Failed", "Couldn't create resource in azure")
		instance.Status.Provisioning = false
		errUpdate := r.KubeClient.Update(ctx, instance)
		if errUpdate != nil {
			//log error and kill it
			r.Recorder.Event(instance, "Warning", "Failed", "Unable to update instance")
		}
		return err
	}
	// write information back to instance
	instance.Status.Provisioning = false
	instance.Status.Provisioned = true
	instance.Status.ProvisionState = azurev1alpha1.Provisioning

	err = r.KubeClient.Update(ctx, instance)
	if err != nil {
		//log error and kill it
		r.Recorder.Event(instance, "Warning", "Failed", "Unable to update instance")
	}

	r.Recorder.Event(instance, "Normal", "Updated", resourcegroupName+" provisioned")

	return nil
}

func (r *AzureController) deleteResourceGroup(definition Definition) error {
	ctx := context.Background()

	var err error
	err = r.ResourceClient.Delete(ctx, definition.Spec)
	if err != nil {
		r.Recorder.Event(definition.Spec, "Warning", "Failed", "Couldn't delete resource in azure")
		return err
	}
	return nil
}

func (r *AzureController) addFinalizer(definition Definition, finalizerName string) error {
	definition.AddFinalizer(finalizerName)
	err := r.KubeClient.Update(context.Background(), definition.Spec)
	if err != nil {
		return fmt.Errorf("failed to update finalizer: %v", err)
	}
	r.Recorder.Event(definition.Spec, "Normal", "Updated", fmt.Sprintf("finalizer %s added", finalizerName))
	return nil
}

func (r *AzureController) handleFinalizer(definition Definition, finalizerName string) error {
	if definition.HasFinalizer(finalizerName) {
		ctx := context.Background()
		// our finalizer is present, so lets handle our external dependency
		if err := r.ResourceClient.Delete(ctx, definition.Spec); err != nil {
			return err
		}

		definition.RemoveFinalizer(r.FinalizerName)
		if err := r.KubeClient.Update(ctx, definition.Spec); err != nil {
			return err
		}
	}
	// Our finalizer has finished, so the reconciler can do nothing.
	return nil
}


func (r *AzureController) deleteResource(instance *azurev1alpha1.ResourceGroup) error {
	ctx := context.Background()

	var err error
	err = r.ResourceClient.Delete(ctx, instance)
	if err != nil {
		r.Recorder.Event(instance, "Warning", "Failed", "Couldn't delete resource in azure")
		return err
	}
	return nil
}
