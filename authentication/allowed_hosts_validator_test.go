package authentication

import (
	u "net/url"
	"testing"

	assert "github.com/stretchr/testify/assert"
)

func TestItValidatesHostsUseNewAllowedHostsValidator(t *testing.T) {
	validator := NewAllowedHostsValidator([]string{"graph.microsoft.com"})
	url, err := u.Parse("https://graph.microsoft.com/v1.0/me")
	assert.Nil(t, err)
	assert.True(t, validator.IsUrlHostValid(url))
}

func TestItValidatesHostsUseNewAllowedHostsValidatorErrorCheck(t *testing.T) {
	_, err := NewAllowedHostsValidatorErrorCheck([]string{"http://graph.microsoft.com"})
	assert.EqualValues(t, ErrInvalidHostPrefix, err)
}
