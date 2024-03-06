package serialization

import "errors"

// UntypedObject defines an untyped object.
type UntypedObject struct {
	UntypedNode
}

// GetValue gets a map of the properties of the object.
func (un *UntypedObject) GetValue() map[string]UntypedNodeable {
	return un.value.(map[string]UntypedNodeable)
}

// NewUntypedObject creates a new UntypedObject object.
func NewUntypedObject(properties map[string]UntypedNodeable) (*UntypedObject, error) {
	if properties == nil {
		return nil, errors.New("properties cannot be nil")
	}
	m := &UntypedObject{}
	m.value = properties
	return m, nil
}
