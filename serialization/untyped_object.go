package serialization

// UntypedObject defines an untyped object.
type UntypedObject struct {
	UntypedNode
}

// GetValue gets a map of the properties of the object.
func (un *UntypedObject) GetValue() map[string]UntypedNodeable {
	return un.value.(map[string]UntypedNodeable)
}

// NewUntypedObject creates a new UntypedObject object.
func NewUntypedObject(properties map[string]UntypedNodeable) *UntypedObject {
	m := &UntypedObject{}
	m.value = properties
	return m
}
