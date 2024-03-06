package serialization

import "errors"

// UntypedFloat defines an untyped float32 value.
type UntypedFloat struct {
	UntypedNode
}

// GetValue returns the float32 value.
func (un *UntypedFloat) GetValue() *float32 {
	return un.value.(*float32)
}

// NewUntypedFloat creates a new UntypedFloat object.
func NewUntypedFloat(float32Value *float32) (*UntypedFloat, error) {
	if float32Value == nil {
		return nil, errors.New("float32Value cannot be nil")
	}
	m := &UntypedFloat{}
	m.value = float32Value
	return m, nil
}
