package serialization

import (
	"testing"

	assert "github.com/stretchr/testify/assert"
)

func TestUntypedNode(t *testing.T) {
	node := NewUntypedNode("sample value")
	assert.NotNil(t, node)
	assert.True(t, node.GetIsUntypedNode())
	assert.Equal(t, "sample value", node.GetValue())
	assert.NotNil(t, node.GetFieldDeserializers())
	assert.Empty(t, node.GetFieldDeserializers())
	assert.NoError(t, node.Serialize(nil))

	discNode, err := CreateUntypedNodeFromDiscriminatorValue(nil)
	assert.NoError(t, err)
	assert.NotNil(t, discNode)
	untypedDiscNode, ok := discNode.(*UntypedNode)
	assert.True(t, ok)
	assert.True(t, untypedDiscNode.GetIsUntypedNode())
	assert.Nil(t, untypedDiscNode.GetValue())

	var nilNode *UntypedNode
	assert.False(t, nilNode.GetIsUntypedNode())
	assert.Nil(t, nilNode.GetValue())
}

func TestUntypedArray(t *testing.T) {
	item1 := NewUntypedString("first")
	item2 := NewUntypedInteger(2)
	collection := []UntypedNodeable{item1, item2}

	arr := NewUntypedArray(collection)
	assert.NotNil(t, arr)
	assert.True(t, arr.GetIsUntypedNode())
	val := arr.GetValue()
	assert.Equal(t, 2, len(val))
	assert.Equal(t, item1, val[0])
	assert.Equal(t, item2, val[1])

	emptyArr := NewUntypedArray(nil)
	assert.NotNil(t, emptyArr)
	assert.Nil(t, emptyArr.GetValue())

	zeroArr := &UntypedArray{}
	assert.Nil(t, zeroArr.GetValue())

	invalidArr := &UntypedArray{}
	invalidArr.value = "not a slice"
	assert.Nil(t, invalidArr.GetValue())

	var nilArr *UntypedArray
	assert.Nil(t, nilArr.GetValue())
}

func TestUntypedBoolean(t *testing.T) {
	bTrue := NewUntypedBoolean(true)
	assert.NotNil(t, bTrue)
	assert.True(t, bTrue.GetIsUntypedNode())
	assert.NotNil(t, bTrue.GetValue())
	assert.True(t, *bTrue.GetValue())

	bFalse := NewUntypedBoolean(false)
	assert.NotNil(t, bFalse)
	assert.NotNil(t, bFalse.GetValue())
	assert.False(t, *bFalse.GetValue())

	zeroBool := &UntypedBoolean{}
	assert.Nil(t, zeroBool.GetValue())

	var nilBool *UntypedBoolean
	assert.Nil(t, nilBool.GetValue())
}

func TestUntypedDouble(t *testing.T) {
	d := NewUntypedDouble(3.14159)
	assert.NotNil(t, d)
	assert.True(t, d.GetIsUntypedNode())
	assert.NotNil(t, d.GetValue())
	assert.Equal(t, 3.14159, *d.GetValue())

	zeroDouble := &UntypedDouble{}
	assert.Nil(t, zeroDouble.GetValue())

	var nilDouble *UntypedDouble
	assert.Nil(t, nilDouble.GetValue())
}

func TestUntypedFloat(t *testing.T) {
	f := NewUntypedFloat(float32(1.23))
	assert.NotNil(t, f)
	assert.True(t, f.GetIsUntypedNode())
	assert.NotNil(t, f.GetValue())
	assert.Equal(t, float32(1.23), *f.GetValue())

	zeroFloat := &UntypedFloat{}
	assert.Nil(t, zeroFloat.GetValue())

	var nilFloat *UntypedFloat
	assert.Nil(t, nilFloat.GetValue())
}

func TestUntypedInteger(t *testing.T) {
	i := NewUntypedInteger(42)
	assert.NotNil(t, i)
	assert.True(t, i.GetIsUntypedNode())
	assert.NotNil(t, i.GetValue())
	assert.Equal(t, int32(42), *i.GetValue())

	zeroInt := &UntypedInteger{}
	assert.Nil(t, zeroInt.GetValue())

	var nilInt *UntypedInteger
	assert.Nil(t, nilInt.GetValue())
}

func TestUntypedLong(t *testing.T) {
	l := NewUntypedLong(9007199254740991)
	assert.NotNil(t, l)
	assert.True(t, l.GetIsUntypedNode())
	assert.NotNil(t, l.GetValue())
	assert.Equal(t, int64(9007199254740991), *l.GetValue())

	zeroLong := &UntypedLong{}
	assert.Nil(t, zeroLong.GetValue())

	var nilLong *UntypedLong
	assert.Nil(t, nilLong.GetValue())
}

func TestUntypedNull(t *testing.T) {
	n := NewUntypedNull()
	assert.NotNil(t, n)
	assert.True(t, n.GetIsUntypedNode())
	assert.Nil(t, n.GetValue())

	zeroNull := &UntypedNull{}
	assert.Nil(t, zeroNull.GetValue())

	var nilNull *UntypedNull
	assert.Nil(t, nilNull.GetValue())
}

func TestUntypedObject(t *testing.T) {
	props := map[string]UntypedNodeable{
		"title": NewUntypedString("Test Object"),
		"count": NewUntypedInteger(10),
	}
	obj := NewUntypedObject(props)
	assert.NotNil(t, obj)
	assert.True(t, obj.GetIsUntypedNode())
	val := obj.GetValue()
	assert.NotNil(t, val)
	assert.Equal(t, 2, len(val))
	assert.Equal(t, props["title"], val["title"])
	assert.Equal(t, props["count"], val["count"])

	zeroObj := &UntypedObject{}
	assert.Nil(t, zeroObj.GetValue())

	invalidObj := &UntypedObject{}
	invalidObj.value = "not a map"
	assert.Nil(t, invalidObj.GetValue())

	var nilObj *UntypedObject
	assert.Nil(t, nilObj.GetValue())
}

func TestUntypedString(t *testing.T) {
	s := NewUntypedString("hello world")
	assert.NotNil(t, s)
	assert.True(t, s.GetIsUntypedNode())
	assert.NotNil(t, s.GetValue())
	assert.Equal(t, "hello world", *s.GetValue())

	zeroStr := &UntypedString{}
	assert.Nil(t, zeroStr.GetValue())

	var nilStr *UntypedString
	assert.Nil(t, nilStr.GetValue())
}

func TestUntypedNodeableInterfaceCompliance(t *testing.T) {
	var _ UntypedNodeable = NewUntypedNode(nil)
	var _ UntypedNodeable = NewUntypedArray(nil)
	var _ UntypedNodeable = NewUntypedBoolean(true)
	var _ UntypedNodeable = NewUntypedDouble(1.0)
	var _ UntypedNodeable = NewUntypedFloat(1.0)
	var _ UntypedNodeable = NewUntypedInteger(1)
	var _ UntypedNodeable = NewUntypedLong(1)
	var _ UntypedNodeable = NewUntypedNull()
	var _ UntypedNodeable = NewUntypedObject(nil)
	var _ UntypedNodeable = NewUntypedString("")

	var _ Parsable = NewUntypedNode(nil)
	var _ Parsable = NewUntypedArray(nil)
	var _ Parsable = NewUntypedBoolean(true)
	var _ Parsable = NewUntypedDouble(1.0)
	var _ Parsable = NewUntypedFloat(1.0)
	var _ Parsable = NewUntypedInteger(1)
	var _ Parsable = NewUntypedLong(1)
	var _ Parsable = NewUntypedNull()
	var _ Parsable = NewUntypedObject(nil)
	var _ Parsable = NewUntypedString("")
}
