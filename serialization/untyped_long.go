package serialization

// UntypedLong defines an untyped int64 value.
type UntypedLong struct {
	UntypedNode
}

// GetValue returns the int64 value.
func (un *UntypedLong) GetValue() *int64 {
	return un.value.(*int64)
}

// NewUntypedLong creates a new UntypedLong object.
func NewUntypedLong(int64Value *int64) *UntypedLong {
	m := &UntypedLong{}
	m.value = int64Value
	return m
}
