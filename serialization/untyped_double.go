package serialization

import "errors"

// UntypedDouble defines an untyped float64 object.
type UntypedDouble struct {
	UntypedNode
}

// GetValue returns the float64 value.
func (un *UntypedDouble) GetValue() *float64 {
	return un.value.(*float64)
}

// NewUntypedDouble creates a new UntypedDouble object.
func NewUntypedDouble(float64Value *float64) (*UntypedDouble, error) {
	if float64Value == nil {
		return nil, errors.New("float64Value cannot be nil")
	}
	m := &UntypedDouble{}
	m.value = float64Value
	return m, nil
}
