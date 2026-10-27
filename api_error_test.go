package abstractions

import (
	"testing"

	assert "github.com/stretchr/testify/assert"
)

func TestApiError(t *testing.T) {
	err := NewApiError()
	assert.NotNil(t, err)
	assert.Equal(t, "error status code received from the API", err.Error())
	assert.NotNil(t, err.GetResponseHeaders())
	assert.Equal(t, 0, err.GetStatusCode())

	err.Message = "custom error message"
	assert.Equal(t, "custom error message", err.Error())

	err.SetStatusCode(404)
	assert.Equal(t, 404, err.GetStatusCode())

	headers := NewResponseHeaders()
	headers.Add("retry-after", "120")
	err.SetResponseHeaders(headers)
	assert.True(t, err.GetResponseHeaders().ContainsKey("retry-after"))
}

func TestApiErrorNilReceiver(t *testing.T) {
	var nilErr *ApiError
	assert.Equal(t, "", nilErr.Error())
	assert.Nil(t, nilErr.GetResponseHeaders())
	assert.Equal(t, 0, nilErr.GetStatusCode())

	headers := NewResponseHeaders()
	nilErr.SetResponseHeaders(headers)
	nilErr.SetStatusCode(500)
}
