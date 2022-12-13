package authentication

import (
	"context"
	"testing"

	abstractions "github.com/microsoft/kiota-abstractions-go"
	assert "github.com/stretchr/testify/assert"
)

func TestItImplementsInterface(t *testing.T) {
	res1, res2 := NewApiKeyAuthenticationProvider("key", "param", HEADER_KEYLOCATION)
	assert.Nil(t, res2)
	assert.NotNil(t, res1)
	res := AuthenticationProvider(res1)
	assert.NotNil(t, res)
}

func TestDefensivePrograming(t *testing.T) {
	res1, res2 := NewApiKeyAuthenticationProvider("", "param", QUERYPARAMETER_KEYLOCATION)
	assert.NotNil(t, res2)
	assert.Nil(t, res1)
	res1, res2 = NewApiKeyAuthenticationProvider("key", "", QUERYPARAMETER_KEYLOCATION)
	assert.NotNil(t, res2)
	assert.Nil(t, res1)

	res1, res2 = NewApiKeyAuthenticationProvider("key", "param", QUERYPARAMETER_KEYLOCATION)
	assert.Nil(t, res2)
	assert.NotNil(t, res1)
	err := res1.AuthenticateRequest(context.Background(), nil, nil)
	assert.NotNil(t, err)
}

func TestItAddsInQueryParameters(t *testing.T) {
	res1, res2 := NewApiKeyAuthenticationProvider("key", "param", QUERYPARAMETER_KEYLOCATION)
	assert.Nil(t, res2)
	assert.NotNil(t, res1)
	request := abstractions.NewRequestInformation()
	request.UrlTemplate = "https://localhost{?param1}"
	err := res1.AuthenticateRequest(context.Background(), request, nil)
	assert.Nil(t, err)
	resultUri, err := request.GetUri()
	assert.Nil(t, err)
	assert.Equal(t, "https://localhost?param=key", resultUri.String())
	assert.False(t, request.Headers.ContainsKey("param"))
}

func TestItAddsInQueryParametersWithOtherParameters(t *testing.T) {
	res1, res2 := NewApiKeyAuthenticationProvider("key", "param", QUERYPARAMETER_KEYLOCATION)
	assert.Nil(t, res2)
	assert.NotNil(t, res1)
	request := abstractions.NewRequestInformation()
	request.UrlTemplate = "https://localhost{?param1}"
	request.QueryParameters["param1"] = "value1"
	err := res1.AuthenticateRequest(context.Background(), request, nil)
	assert.Nil(t, err)
	resultUri, err := request.GetUri()
	assert.Nil(t, err)
	assert.Equal(t, "https://localhost?param=key&param1=value1", resultUri.String())
	assert.False(t, request.Headers.ContainsKey("param"))
}

func TestItAddsInHeader(t *testing.T) {
	res1, res2 := NewApiKeyAuthenticationProvider("key", "param", HEADER_KEYLOCATION)
	assert.Nil(t, res2)
	assert.NotNil(t, res1)
	request := abstractions.NewRequestInformation()
	request.UrlTemplate = "https://localhost{?param1}"
	err := res1.AuthenticateRequest(context.Background(), request, nil)
	assert.Nil(t, err)
	resultUri, err := request.GetUri()
	assert.Nil(t, err)
	assert.Equal(t, "https://localhost", resultUri.String())
	assert.Equal(t, "key", request.Headers.Get("param")[0])
}
