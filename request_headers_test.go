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
}

func TestIdAdds(t *testing.T) {
	instance := NewRequestHeaders()
	assert.NotNil(t, instance)
	instance.Add("key", "value")
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

func TestIdListKeys(t *testing.T) {
	instance := NewRequestHeaders()
	assert.NotNil(t, instance)
	instance.Add("key", "value")
	assert.True(t, instance.ContainsKey("key"))
	assert.Equal(t, "key", instance.ListKeys()[0])
}
