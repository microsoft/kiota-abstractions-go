package serialization

import "errors"

// UntypedArray defines an untyped collection.
type UntypedArray struct {
	UntypedNode
}

// GetValue returns a collection of UntypedNode.
func (un *UntypedArray) GetValue() []UntypedNodeable {
	return un.value.([]UntypedNodeable)
}

// NewUntypedArray creates a new UntypedArray object.
func NewUntypedArray(collection []UntypedNodeable) (*UntypedArray, error) {
	if collection == nil {
		return nil, errors.New("collection cannot be nil")
	}
	m := &UntypedArray{}
	m.value = collection
	return m, nil
}
