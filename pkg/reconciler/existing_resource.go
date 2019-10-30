package reconciler

type ExistingResourceBehaviour string

const (
	ExistingResourceReject ExistingResourceBehaviour = "Reject"
	ExistingResourceView   ExistingResourceBehaviour = "View"
	ExistingResourceManage ExistingResourceBehaviour = "Manage"
)

func (e ExistingResourceBehaviour) reject() bool {
	return e == ExistingResourceReject
}

func (e ExistingResourceBehaviour) view() bool {
	return !e.reject() && !e.manage() // this is the default behaviour
}

func (e ExistingResourceBehaviour) manage() bool {
	return e == ExistingResourceManage
}

const (
	ExistingResourceBehaviourAnnotation = "azure.microsoft.com/existing-resource-behaviour"
	ReadOnlyResourceAnnotation          = "azure.microsoft.com/read-only-resource"
)
