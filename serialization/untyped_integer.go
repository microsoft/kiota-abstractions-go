package serialization

import "errors"

// UntypedInteger defines an untyped integer value.
type UntypedInteger struct {
	UntypedNode
}

// GetValue returns the int32 value.
func (un *UntypedInteger) GetValue() *int32 {
	return un.value.(*int32)
}

// NewUntypedInteger creates a new UntypedInteger object.
func NewUntypedInteger(int32Value *int32) (*UntypedInteger, error) {
	if int32Value == nil {
		return nil, errors.New("int32Value cannot be nil")
	}
	m := &UntypedInteger{}
	m.value = int32Value
	return m, nil
}
