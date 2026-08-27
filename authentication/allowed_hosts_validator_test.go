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

func TestItValidatesSubdomainMatchingAllowedSuffix(t *testing.T) {
	validator := NewAllowedHostsValidator([]string{".fabric.microsoft.com"})
	url, err := u.Parse("https://abc.123.graphql.fabric.microsoft.com/path")
	assert.Nil(t, err)
	assert.True(t, validator.IsUrlHostValid(url))
}

func TestItRejectsBareDomainWhenAllowedAsSuffix(t *testing.T) {
	validator := NewAllowedHostsValidator([]string{".fabric.microsoft.com"})
	url, err := u.Parse("https://fabric.microsoft.com/path")
	assert.Nil(t, err)
	assert.False(t, validator.IsUrlHostValid(url))
}

func TestItValidatesSuffixHostCaseInsensitively(t *testing.T) {
	validator := NewAllowedHostsValidator([]string{".Fabric.Microsoft.COM"})
	url, err := u.Parse("https://ABC.z2c.graphql.fabric.microsoft.com/path")
	assert.Nil(t, err)
	assert.True(t, validator.IsUrlHostValid(url))
}

func TestItAllowsMultipleValidHostsWithSuffixes(t *testing.T) {
	validator := NewAllowedHostsValidator([]string{"example.com", "api.example.com", ".fabric.microsoft.com"})
	testCases := []struct {
		url      string
		expected bool
	}{
		{"https://example.com/path", true},
		{"https://api.example.com/path", true},
		{"https://other.com/path", false},
		{"https://abc.123.graphql.fabric.microsoft.com/path", true},
	}

	for _, testCase := range testCases {
		url, err := u.Parse(testCase.url)
		assert.Nil(t, err)
		assert.Equal(t, testCase.expected, validator.IsUrlHostValid(url))
	}
}

func TestItAllowsSuffixBasedHostsAfterUpdate(t *testing.T) {
	validator := NewAllowedHostsValidator([]string{"example.com"})
	validator.SetAllowedHosts([]string{".fabric.microsoft.com"})
	url, err := u.Parse("https://abc.123.graphql.fabric.microsoft.com/path")
	assert.Nil(t, err)
	assert.True(t, validator.IsUrlHostValid(url))
}

func TestItValidatesHostsUseNewAllowedHostsValidatorErrorCheck(t *testing.T) {
	validator, err := NewAllowedHostsValidatorErrorCheck([]string{"http://graph.microsoft.com"})
	assert.EqualValues(t, ErrInvalidHostPrefix, err)
	assert.Nil(t, validator)
}
