package astmodel

type ReadonlyObjectType interface {
	// Properties returns all the property definitions
	Properties() ReadOnlyPropertySet

	// Property returns the details of a specific property based on its unique case-sensitive name
	Property(name PropertyName) (*PropertyDefinition, bool)

	// EmbeddedProperties returns all the embedded properties
	// A sorted slice is returned to preserve immutability and provide determinism
	EmbeddedProperties() []*PropertyDefinition

	IsResource() bool
	Resources() TypeNameSet

	// Functions returns all the function implementations
	// A sorted slice is returned to preserve immutability and provide determinism
	Functions() []Function

	// HasFunctionWithName determines if this object has a function with the given name
	HasFunctionWithName(name string) bool

	TestCases() []TestCase
}
