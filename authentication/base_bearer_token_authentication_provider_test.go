package authentication

import (
	"context"
	u "net/url"
	"testing"

	assert "github.com/stretchr/testify/assert"
)

type MockAccessTokenProvider struct {
}

func (m *MockAccessTokenProvider) GetAuthorizationToken(ctx context.Context, url *u.URL, additionalAuthenticationContext map[string]interface{}) (string, error) {
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
