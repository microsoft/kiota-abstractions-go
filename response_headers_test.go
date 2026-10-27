package abstractions

import (
	"testing"

	assert "github.com/stretchr/testify/assert"
)

func TestResponseHeadersDefensiveProgramming(t *testing.T) {
	instance := NewResponseHeaders()
	assert.NotNil(t, instance)
	instance.Add("", "")
	assert.False(t, instance.ContainsKey(""))
	instance.Add("key", "")
	assert.False(t, instance.ContainsKey("key"))
	instance.Remove("")
	instance.RemoveValue("", "")
	instance.RemoveValue("key", "")
	instance.AddAll(nil)

	instance = &ResponseHeaders{}
	assert.NotNil(t, instance)
	instance.Get("key")
	instance.Remove("key")
	instance.ListKeys()
	instance.RemoveValue("key", "value")
	instance.Add("key", "value")

	var nilInstance *ResponseHeaders
	nilInstance.Add("key", "value")
	assert.Nil(t, nilInstance.Get("key"))
	nilInstance.Remove("key")
	nilInstance.RemoveValue("key", "value")
	assert.False(t, nilInstance.ContainsKey("key"))
	nilInstance.Clear()
	assert.Empty(t, nilInstance.ListKeys())
	nilInstance.AddAll(instance)
}

func TestResponseHeadersAdds(t *testing.T) {
	instance := NewResponseHeaders()
	assert.NotNil(t, instance)
	instance.Add("key", "value")
	assert.True(t, instance.ContainsKey("key"))
	assert.Equal(t, "value", instance.Get("key")[0])
}

func TestResponseHeadersRemoves(t *testing.T) {
	instance := NewResponseHeaders()
	assert.NotNil(t, instance)
	instance.Add("key", "value")
	assert.True(t, instance.ContainsKey("key"))
	instance.Remove("key")
	assert.False(t, instance.ContainsKey("key"))
}

func TestResponseHeadersRemovesValue(t *testing.T) {
	instance := NewResponseHeaders()
	assert.NotNil(t, instance)
	instance.Add("key", "value", "value2")
	assert.True(t, instance.ContainsKey("key"))
	instance.RemoveValue("key", "value")
	assert.True(t, instance.ContainsKey("key"))
	assert.Equal(t, "value2", instance.Get("key")[0])
	instance.RemoveValue("key", "value2")
	assert.False(t, instance.ContainsKey("key"))
}

func TestResponseHeadersClears(t *testing.T) {
	instance := NewResponseHeaders()
	assert.NotNil(t, instance)
	instance.Add("key", "value")
	assert.True(t, instance.ContainsKey("key"))
	instance.Clear()
	assert.False(t, instance.ContainsKey("key"))
}

func TestResponseHeadersAddAll(t *testing.T) {
	instance1 := NewResponseHeaders()
	instance1.Add("key1", "val1")

	instance2 := NewResponseHeaders()
	instance2.Add("key2", "val2")
	instance2.Add("key1", "val1_2")

	instance1.AddAll(instance2)
	assert.True(t, instance1.ContainsKey("key1"))
	assert.True(t, instance1.ContainsKey("key2"))
	assert.Equal(t, 2, len(instance1.Get("key1")))
}

func TestResponseHeadersTryAdd(t *testing.T) {
	instance := NewResponseHeaders()
	assert.NotNil(t, instance)
	assert.True(t, instance.TryAdd("key", "value"))
	assert.False(t, instance.TryAdd("key", "value2"))
	assert.True(t, instance.ContainsKey("key"))
	assert.Equal(t, "value", instance.Get("key")[0])

	assert.False(t, instance.TryAdd("anotherKey", ""))
	assert.False(t, instance.ContainsKey("anotherKey"))
	assert.False(t, instance.TryAdd("   ", "val"))
	assert.False(t, instance.ContainsKey("   "))
	assert.False(t, instance.TryAdd("", "val"))

	var nilInstance *ResponseHeaders
	assert.False(t, nilInstance.TryAdd("key", "value"))
}

