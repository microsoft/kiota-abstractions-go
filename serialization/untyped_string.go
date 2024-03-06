package serialization

import "errors"

// UntypedString defines an untyped string object.
type UntypedString struct {
	UntypedNode
}

// GetValue returns the string object.
func (un *UntypedString) GetValue() *string {
	return un.value.(*string)
}

// NewUntypedString creates a new UntypedString object.
func NewUntypedString(stringValue *string) (*UntypedString, error) {
	if stringValue == nil {
		return nil, errors.New("stringValue cannot be nil")
	}
	m := &UntypedString{}
	m.value = stringValue
	return m, nil
}
