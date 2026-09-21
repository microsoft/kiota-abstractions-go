package serialization

import "math/big"

// UntypedFloat128 defines an untyped float128 object.
type UntypedFloat128 struct {
	UntypedNode
}

// GetValue returns the *big.Float value.
func (un *UntypedFloat128) GetValue() *big.Float {
	castValue, ok := un.value.(*big.Float)
	if ok {
		return castValue
	}
	return nil
}

// NewUntypedFloat128 creates a new UntypedFloat128 object.
func NewUntypedFloat128(float128Value *big.Float) *UntypedFloat128 {
	m := &UntypedFloat128{}
	m.value = float128Value
	return m
}
