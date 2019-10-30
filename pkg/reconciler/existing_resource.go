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
	ExistingResourceBehaviourAnnotation = "/existing-resource-behaviour"
	ReadOnlyResourceAnnotation          = "/read-only-resource"
)
