package v1

import (
	resoucegroupsresourcemanager "github.com/Azure/azure-service-operator/pkg/resourcemanager/resourcegroups"
	"golang.org/x/net/context"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

var eventhubnamespacelog = logf.Log.WithName("eventhubnamespace-resource")

func (eventhubNamespace *EventhubNamespace) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(eventhubNamespace).
		Complete()
}

// +kubebuilder:webhook:path=/mutate-azure-microsoft-com-v1-eventhubnamespace,mutating=true,failurePolicy=fail,groups=azure.microsoft.com,resources=eventhubnamespaces,verbs=create;update,versions=v1,name=meventhubnamespace.kb.io

var _ webhook.Defaulter = &EventhubNamespace{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (eventhubNamespace *EventhubNamespace) Default() {
	eventhubnamespacelog.Info("default", "name", eventhubNamespace.Name)

	// place to add defaulting logic.
}

// +kubebuilder:webhook:path=/validate-azure-microsoft-com-v1-eventhubnamespace,mutating=false,failurePolicy=fail,groups=azure.microsoft.com,resources=eventhubnamespaces,verbs=create;update,versions=v1,name=veventhubnamespace.kb.io

var _ webhook.Validator = &EventhubNamespace{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (eventhubNamespace *EventhubNamespace) ValidateCreate() error {
	eventhubnamespacelog.Info("validate create", "name", eventhubNamespace.Name)

	//validation logic upon object creation.
	return eventhubNamespace.validateEventhubNamespace()
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (eventhubNamespace *EventhubNamespace) ValidateUpdate(old runtime.Object) error {
	eventhubnamespacelog.Info("validate update", "name", eventhubNamespace.Name)

	//validation logic upon object update.
	return eventhubNamespace.validateEventhubNamespace()
}

func (eventhubNamespace *EventhubNamespace) validateEventhubNamespace() error {
	var allErrs field.ErrorList

	if len(eventhubNamespace.Name) < 6 {
		err := field.Invalid(field.NewPath("metadata").Child("name"), eventhubNamespace.Name, "Eventhubnamespace minimum length should be 6")
		allErrs = append(allErrs, err)
	}

	err := eventhubNamespace.validateValidParent()
	if err != nil {
		eventhubnamespacelog.Info("error", "name", eventhubNamespace.Name)
		allErrs = append(allErrs, err)
	}
	if len(allErrs) == 0 {
		return nil
	}

	return apierrors.NewInvalid(
		schema.GroupKind{Group: "azure.microsoft.com", Kind: "EventhubNamespace"},
		eventhubNamespace.Name, allErrs)
}

func (eventhubNamespace *EventhubNamespace) validateValidParent() *field.Error {

	resourcegroupName := eventhubNamespace.Spec.ResourceGroup
	eventhubnamespacelog.Info("validate resource group Name", "resourcegroupName", resourcegroupName)
	result, _ := resoucegroupsresourcemanager.CheckExistence(context.Background(), resourcegroupName)
	eventhubnamespacelog.Info("validate resource group Name", "result.Response.StatusCode", result.Response.StatusCode)
	if result.Response.StatusCode != 204 {
		return field.Invalid(field.NewPath("spec").Child("resourcegroup"), resourcegroupName, "ResourceGroup doesn't exist")
	}
	return nil
}
