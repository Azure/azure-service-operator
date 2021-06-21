package astmodel

// TypeNameProcessingRegister is a register of TypeNames that require processing, keeping track of which names have been
// processed and which are still outstanding. This is used, for example, for a consistency check that all TypeName
// references can be satisfied
type TypeNameProcessingRegister struct {
	// known is the set of all known type names, each with reasons why it's included
	known map[TypeName][]string
	// completed is the set of all names that have been processed
	completed TypeNameSet
}

// NewTypeNameProcessingRegister constructs a new empty TypeNameProcessingRegister
func NewTypeNameProcessingRegister() *TypeNameProcessingRegister {
	return &TypeNameProcessingRegister{
		completed: NewTypeNameSet(),
		known:     make(map[TypeName][]string),
	}
}

// IsCompleted returns true if the specified TypeName is in the register and processing has been completed
func (reg *TypeNameProcessingRegister) IsCompleted(name TypeName) bool {
	_, found := reg.known[name]
	return found && reg.completed.Contains(name)
}

// IsPending returns true if the specified TypeName is in the register and processing has not been completed
func (reg *TypeNameProcessingRegister) IsPending(name TypeName) bool {
	_, found := reg.known[name]
	return found && !reg.completed.Contains(name)
}

// Complete marks the specified name as processing complete
// As a side effect, the name will be added to the register if it is missing
func (reg *TypeNameProcessingRegister) Complete(name TypeName) {
	if _, found := reg.known[name]; !found {
		reg.known[name] = []string{}
	}

	reg.completed.Add(name)
}

// Reset marks the specified name as unprocessed (ie not complete)
// As a side effect, the name will be added to the checklist if it is missing
func (reg *TypeNameProcessingRegister) Reset(name TypeName) {
	if _, found := reg.known[name]; !found {
		reg.known[name] = []string{}
	}

	reg.completed.Remove(name)
}

// Requires indicates that processing of the specified name is required - this may already have taken place.
// name is the TypeName to include
// because is the reason why it has been included
// Other reasons are preserved, if present
func (reg *TypeNameProcessingRegister) Requires(name TypeName, because string) {
	reg.known[name] = append(reg.known[name], because)
}

// Remove removes the specified name from the list
// A single call to Remove can undo many calls to Requires()
func (reg *TypeNameProcessingRegister) Remove(name TypeName) {
	delete(reg.known, name)
	reg.completed.Remove(name)
}

// Length returns the count of names in the checklist
func (reg *TypeNameProcessingRegister) Length() int {
	return len(reg.known)
}

// Pending gets a map all of the type names that have not been processed, along with the reasons why they were included
// in the register in the first place
func (reg *TypeNameProcessingRegister) Pending() map[TypeName][]string {
	return reg.collect(func(name TypeName) bool {
		_, completed := reg.completed[name]
		return !completed
	})
}

// Completed gets all of the type names that have been processed, along with the reasons why they were included in the
// register in the first place
func (reg *TypeNameProcessingRegister) Completed() map[TypeName][]string {
	return reg.collect(func(name TypeName) bool {
		_, completed := reg.completed[name]
		return completed
	})
}

func (reg *TypeNameProcessingRegister) collect(predicate func(TypeName) bool) map[TypeName][]string {
	result := make(map[TypeName][]string)
	for name, because := range reg.known {
		if predicate(name) {
			result[name] = because
		}
	}

	return result
}
