package abstractions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBaseRequestBuilderWithTypedNilAdapter(t *testing.T) {
	var adapter *MockRequestAdapter
	assert.PanicsWithValue(t, "requestAdapter cannot be nil", func() {
		NewBaseRequestBuilder(adapter, "https://example.com", nil)
	})
}

func TestNewBaseRequestBuilderSuccess(t *testing.T) {
	adapter := &MockRequestAdapter{}
	params := map[string]string{"key": "value"}
	builder := NewBaseRequestBuilder(adapter, "https://example.com/{key}", params)
	assert.NotNil(t, builder)
	assert.Equal(t, adapter, builder.RequestAdapter)
	assert.Equal(t, "https://example.com/{key}", builder.UrlTemplate)
	assert.Equal(t, "value", builder.PathParameters["key"])
}
