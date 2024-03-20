package serialization

import "errors"

// UntypedLong defines an untyped int64 value.
type UntypedLong struct {
	UntypedNode
}

// GetValue returns the int64 value.
func (un *UntypedLong) GetValue() *int64 {
	castValue, ok := un.value.(*int64)
	if ok {
		return castValue
	}
	return nil
}

// NewUntypedLong creates a new UntypedLong object.
func NewUntypedLong(int64Value *int64) (*UntypedLong, error) {
	if int64Value == nil {
		return nil, errors.New("int64Value cannot be nil")
	}
	m := &UntypedLong{}
	m.value = int64Value
	return m, nil
}
