package authentication

import (
	u "net/url"
	"testing"

	assert "github.com/stretchr/testify/assert"
)

type MockAccessTokenProvider struct {
}

func (m *MockAccessTokenProvider) GetAuthorizationToken(url *u.URL, additionalAuthenticationContext map[string]interface{}) (string, error) {
	return "", nil
}
func (m *MockAccessTokenProvider) GetAllowedHostsValidator() *AllowedHostsValidator {
	return nil
}
func TestBaseBearerProviderHonoursInterface(t *testing.T) {
	mockToken := &MockAccessTokenProvider{}
	instance := NewBaseBearerTokenAuthenticationProvider(mockToken)
	assert.Implements(t, (*AuthenticationProvider)(nil), instance)
}
