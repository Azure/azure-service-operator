package v1

import (
	eventhubsmanager "github.com/Azure/azure-service-operator/pkg/resourcemanager/eventhubs"
	"golang.org/x/net/context"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

var eventhublog = logf.Log.WithName("eventhub-resource")

//SetupWebhookWithManager add webhook manager
func (eventhub *Eventhub) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(eventhub).
		Complete()
}

// +kubebuilder:webhook:path=/mutate-azure-microsoft-com-v1-eventhub,mutating=true,failurePolicy=fail,groups=azure.microsoft.com,resources=eventhubs,verbs=create;update,versions=v1,name=meventhub.kb.io

var _ webhook.Defaulter = &Eventhub{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (eventhub *Eventhub) Default() {
	eventhublog.Info("default", "name", eventhub.Name)

	// place to add defaulting logic.
}

// +kubebuilder:webhook:path=/validate-azure-microsoft-com-v1-eventhub,mutating=false,failurePolicy=fail,groups=azure.microsoft.com,resources=eventhubs,verbs=create;update,versions=v1,name=veventhub.kb.io

var _ webhook.Validator = &Eventhub{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (eventhub *Eventhub) ValidateCreate() error {
	eventhublog.Info("validate create", "name", eventhub.Name)

	//validation logic upon object creation.
	return eventhub.validateeventhub()
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (eventhub *Eventhub) ValidateUpdate(old runtime.Object) error {
	eventhublog.Info("validate update", "name", eventhub.Name)

	//validation logic upon object update.
	return eventhub.validateeventhub()
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (eventhub *Eventhub) ValidateDelete() error {
	eventhublog.Info("validate delete", "name", eventhub.Name)

	//validation logic upon object deletion.
	return nil
}

func (eventhub *Eventhub) validateeventhub() error {
	var allErrs field.ErrorList
	err := eventhub.validateValidParent()
	if err != nil {
		eventhublog.Info("error", "name", eventhub.Name)
		allErrs = append(allErrs, err)
	}
	if len(allErrs) == 0 {
		return nil
	}

	return apierrors.NewInvalid(
		schema.GroupKind{Group: "azure.microsoft.com", Kind: "eventhub"},
		eventhub.Name, allErrs)
}

func (eventhub *Eventhub) validateValidParent() *field.Error {

	resourcegroupName := eventhub.Spec.ResourceGroup
	eventhubNamespaceName := eventhub.Spec.Namespace
	eventhublog.Info("validate event hub namespace Name", "eventhubNamespaceName", eventhubNamespaceName)
	result, _ := eventhubsmanager.GetNamespace(context.Background(), resourcegroupName, eventhubNamespaceName)
	eventhublog.Info("validate event hub namespace Name", "result.Response.StatusCode", result.Response.StatusCode)
	if result.Response.StatusCode != 200 {
		return field.Invalid(field.NewPath("spec").Child("namespace"), eventhubNamespaceName, "EventHubNamespace doesn't exist")
	}
	return nil
}
