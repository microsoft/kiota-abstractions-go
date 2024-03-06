package serialization

import "errors"

// UntypedBoolean defines an untyped boolean object.
type UntypedBoolean struct {
	UntypedNode
}

// GetValue returns the bool value.
func (un *UntypedBoolean) GetValue() *bool {
	return un.value.(*bool)
}

// NewUntypedBoolean creates a new UntypedBoolean object.
func NewUntypedBoolean(boolValue *bool) (*UntypedBoolean, error) {
	if boolValue == nil {
		return nil, errors.New("boolValue cannot be nil")
	}
	m := &UntypedBoolean{}
	m.value = boolValue
	return m, nil
}
