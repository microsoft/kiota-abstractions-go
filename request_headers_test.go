package abstractions

import (
	"testing"

	assert "github.com/stretchr/testify/assert"
)

func TestItFollowsDefensivePrograming(t *testing.T) {
	instance := NewRequestHeaders()
	assert.NotNil(t, instance)
	instance.Add("", "")
	assert.False(t, instance.ContainsKey(""))
	instance.Add("key", "")
	assert.False(t, instance.ContainsKey("key"))
	instance.Remove("")
	instance.RemoveValue("", "")
	instance.RemoveValue("key", "")
	instance.AddAll(nil)

	instance = &RequestHeaders{}
	assert.NotNil(t, instance)
	instance.Get("key")
	instance.Remove("key")
	instance.ListKeys()
	instance.RemoveValue("key", "value")
	instance.Add("key", "value")

}

func TestIdAdds(t *testing.T) {
	instance := NewRequestHeaders()
	assert.NotNil(t, instance)
	instance.Add("key", "value")
	assert.True(t, instance.ContainsKey("key"))
	assert.Equal(t, "value", instance.Get("key")[0])
}

func TestIdTryAdds(t *testing.T) {
	instance := NewRequestHeaders()
	assert.NotNil(t, instance)
	assert.True(t, instance.TryAdd("key", "value"))
	assert.True(t, instance.ContainsKey("key"))
	assert.Equal(t, "value", instance.Get("key")[0])
}

func TestItRemoves(t *testing.T) {
	instance := NewRequestHeaders()
	assert.NotNil(t, instance)
	instance.Add("key", "value")
	assert.True(t, instance.ContainsKey("key"))
	instance.Remove("key")
	assert.False(t, instance.ContainsKey("key"))
}
func TestItRemovesValue(t *testing.T) {
	instance := NewRequestHeaders()
	assert.NotNil(t, instance)
	instance.Add("key", "value", "value2")
	assert.True(t, instance.ContainsKey("key"))
	instance.RemoveValue("key", "value")
	assert.True(t, instance.ContainsKey("key"))
	assert.Equal(t, "value2", instance.Get("key")[0])
}
func TestItClears(t *testing.T) {
	instance := NewRequestHeaders()
	assert.NotNil(t, instance)
	instance.Add("key", "value")
	assert.True(t, instance.ContainsKey("key"))
	instance.Clear()
	assert.False(t, instance.ContainsKey("key"))
}
func TestItAddsAll(t *testing.T) {
	instance := NewRequestHeaders()
	assert.NotNil(t, instance)
	instance2 := NewRequestHeaders()
	instance2.Add("key", "value")
	assert.False(t, instance.ContainsKey("key"))
	instance.AddAll(instance2)
	assert.True(t, instance.ContainsKey("key"))
	assert.Equal(t, "value", instance.Get("key")[0])
}

func TestItTryAdds(t *testing.T) {
	instance := NewRequestHeaders()
	assert.NotNil(t, instance)
	assert.True(t, instance.TryAdd("key", "value"))
	assert.True(t, instance.ContainsKey("key"))
	assert.False(t, instance.TryAdd("key", "value2"))
	assert.Equal(t, "value", instance.Get("key")[0])
}

func TestIdListKeys(t *testing.T) {
	instance := NewRequestHeaders()
	assert.NotNil(t, instance)
	instance.Add("key", "value")
	assert.True(t, instance.ContainsKey("key"))
	assert.Equal(t, "key", instance.ListKeys()[0])
}

func TestNilReceiverAddAllDoesNotPanic(t *testing.T) {
	var nilHeaders *RequestHeaders
	other := NewRequestHeaders()
	other.Add("key", "value")
	assert.NotPanics(t, func() {
		nilHeaders.AddAll(other)
	})

	var nilResponseHeaders *ResponseHeaders
	otherResp := NewResponseHeaders()
	otherResp.Add("key", "value")
	assert.NotPanics(t, func() {
		nilResponseHeaders.AddAll(otherResp)
	})
}

func TestAddWithEmptyFirstValueAndValidAdditionalValues(t *testing.T) {
	instance := NewRequestHeaders()
	instance.Add("key", "", "val1", "val2")
	assert.True(t, instance.ContainsKey("key"))
	values := instance.Get("key")
	assert.Len(t, values, 2)
	assert.Contains(t, values, "val1")
	assert.Contains(t, values, "val2")
}

func TestHeaderKeyWhitespaceTrimming(t *testing.T) {
	instance := NewRequestHeaders()
	instance.Add("\tcontent-type\r\n", "application/json")
	assert.True(t, instance.ContainsKey("content-type"))
	assert.Equal(t, "application/json", instance.Get("content-type")[0])
}

func TestConfigureRequestInformationWithNilHeaders(t *testing.T) {
	req := &RequestInformation{}
	cfg := &RequestConfiguration[struct{}]{
		Headers: NewRequestHeaders(),
	}
	cfg.Headers.Add("Authorization", "Bearer token")
	assert.NotPanics(t, func() {
		ConfigureRequestInformation(req, cfg)
	})
	assert.NotNil(t, req.Headers)
	assert.True(t, req.Headers.ContainsKey("authorization"))
}

