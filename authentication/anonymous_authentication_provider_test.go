package authentication

import (
	"testing"

	assert "github.com/stretchr/testify/assert"
)

func TestAnonymousProviderHonoursInterface(t *testing.T) {
	instance := &AnonymousAuthenticationProvider{}
	assert.Implements(t, (*AuthenticationProvider)(nil), instance)
}
