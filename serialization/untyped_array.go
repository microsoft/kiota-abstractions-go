package serialization

// UntypedArray defines an untyped collection.
type UntypedArray struct {
	UntypedNode
}

// GetValue returns a collection of UntypedNode.
func (un *UntypedArray) GetValue() []UntypedNodeable {
	if un == nil {
		return nil
	}
	castValue, ok := un.value.([]UntypedNodeable)
	if ok {
		return castValue
	}
	return nil
}

// NewUntypedArray creates a new UntypedArray object.
func NewUntypedArray(collection []UntypedNodeable) *UntypedArray {
	m := &UntypedArray{}
	m.value = collection
	return m
}
