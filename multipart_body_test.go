package abstractions

import (
	"testing"

	"github.com/microsoft/kiota-abstractions-go/internal"
	serialization "github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/stretchr/testify/assert"
)

func TestMultipartIsParsable(t *testing.T) {
	multipart := NewMultipartBody()
	if _, ok := multipart.(serialization.Parsable); !ok {
		t.Errorf("MultipartBody does not implement Parsable")
	}
}

func TestMultipartImplementsDefensiveProgramming(t *testing.T) {
	multipart := NewMultipartBody()
	if err := multipart.AddOrReplacePart("", "foo", "bar"); err == nil {
		t.Errorf("AddOrReplacePart should return an error when name is empty")
	}
	if err := multipart.AddOrReplacePart("foo", "", "bar"); err == nil {
		t.Errorf("AddOrReplacePart should return an error when contentType is empty")
	}
	if err := multipart.AddOrReplacePart("foo", "bar", nil); err == nil {
		t.Errorf("AddOrReplacePart should return an error when content is nil")
	}
	if err := multipart.RemovePart(""); err == nil {
		t.Errorf("RemovePart should return an error when name is empty")
	}
	if _, err := multipart.GetPartValue(""); err == nil {
		t.Errorf("GetPartValue should return an error when name is empty")
	}
	if err := multipart.Serialize(nil); err == nil {
		t.Errorf("Serialize should return an error when writer is nil")
	}
}

func TestItRequiresARequestAdapter(t *testing.T) {
	multipart := NewMultipartBody()
	mockSerializer := &internal.MockSerializer{}
	if err := multipart.Serialize(mockSerializer); err == nil {
		t.Errorf("Serialize should return an error when request adapter is nil")
	}
}

func TestItRequiresParts(t *testing.T) {
	multipart := NewMultipartBody()
	mockSerializer := &internal.MockSerializer{}
	mockRequestAdapter := &MockRequestAdapter{}
	multipart.SetRequestAdapter(mockRequestAdapter)
	if err := multipart.Serialize(mockSerializer); err == nil {
		t.Errorf("Serialize should return an error when request adapter is nil")
	}
}

func TestItAddsAPart(t *testing.T) {
	multipart := NewMultipartBody()
	mockRequestAdapter := &MockRequestAdapter{}
	multipart.SetRequestAdapter(mockRequestAdapter)
	err := multipart.AddOrReplacePart("foo", "bar", "baz")
	assert.Nil(t, err)
	value, err := multipart.GetPartValue("foo")
	assert.Nil(t, err)
	valueString, ok := value.(string)
	assert.True(t, ok)
	assert.Equal(t, "baz", valueString)
}

func TestItRemovesPart(t *testing.T) {
	multipart := NewMultipartBody()
	mockRequestAdapter := &MockRequestAdapter{}
	multipart.SetRequestAdapter(mockRequestAdapter)
	err := multipart.AddOrReplacePart("foo", "bar", "baz")
	assert.Nil(t, err)
	err = multipart.RemovePart("FOO")
	assert.Nil(t, err)
	value, err := multipart.GetPartValue("foo")
	assert.Nil(t, err)
	assert.Nil(t, value)
}

//serialize method is being tested in the serialization library
type errorMockSerializerFactory struct{}

func (*errorMockSerializerFactory) GetValidContentType() (string, error) {
	return "application/json", nil
}

func (*errorMockSerializerFactory) GetSerializationWriter(contentType string) (serialization.SerializationWriter, error) {
	return nil, assert.AnError
}

type dummyParsable struct {
	serialization.Parsable
}

func (*dummyParsable) Serialize(writer serialization.SerializationWriter) error {
	return nil
}

func (*dummyParsable) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return nil
}

func TestMultipartSerializationErrorDoesNotPanicOnClose(t *testing.T) {
	multipart := NewMultipartBody()
	mockWriter := &internal.MockSerializer{
		CallsCounter: make(map[string]int),
	}
	mockRequestAdapter := &MockRequestAdapter{
		SerializationWriterFactory: &errorMockSerializerFactory{},
	}
	multipart.SetRequestAdapter(mockRequestAdapter)

	err := multipart.AddOrReplacePart("part1", "application/json", &dummyParsable{})
	assert.NoError(t, err)

	assert.NotPanics(t, func() {
		err := multipart.Serialize(mockWriter)
		assert.Error(t, err)
	})
}

func TestMultipartBodyNilReceiverGuards(t *testing.T) {
	var nilMultipart *multipartBody
	assert.NotPanics(t, func() {
		assert.Error(t, nilMultipart.AddOrReplacePart("part1", "text/plain", "content"))
		val, err := nilMultipart.GetPartValue("part1")
		assert.Error(t, err)
		assert.Nil(t, val)
		assert.Error(t, nilMultipart.RemovePart("part1"))
		assert.Error(t, nilMultipart.Serialize(nil))
		assert.Nil(t, nilMultipart.GetRequestAdapter())
		nilMultipart.SetRequestAdapter(nil)
		assert.Empty(t, nilMultipart.GetBoundary())
	})
}

func TestMultipartBodyTypedNilSafety(t *testing.T) {
	multipart := NewMultipartBody()
	mockWriter := &internal.MockSerializer{
		CallsCounter: make(map[string]int),
	}

	// 1. typed nil content in AddOrReplacePart
	var nilParsable *dummyParsable
	err := multipart.AddOrReplacePart("part1", "application/json", nilParsable)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "content cannot be nil")

	// 2. typed nil request adapter in Serialize
	var nilAdapter *MockRequestAdapter
	multipart.SetRequestAdapter(nilAdapter)
	err = multipart.Serialize(mockWriter)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "requestAdapter cannot be nil")

	// 3. typed nil serialization writer factory in Serialize
	var nilFactory *errorMockSerializerFactory
	mockAdapterWithNilFactory := &MockRequestAdapter{
		SerializationWriterFactory: nilFactory,
	}
	multipart.SetRequestAdapter(mockAdapterWithNilFactory)
	_ = multipart.AddOrReplacePart("part2", "text/plain", "valid content")
	err = multipart.Serialize(mockWriter)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "serializationWriterFactory cannot be nil")
}


func TestMultipartBodyNilByteSliceSupport(t *testing.T) {
	multipart := NewMultipartBody()

	// 1. var data []byte (nil slice)
	var nilSlice []byte
	err := multipart.AddOrReplacePart("binaryPart1", "application/octet-stream", nilSlice)
	assert.NoError(t, err)

	val, err := multipart.GetPartValue("binaryPart1")
	assert.NoError(t, err)
	assert.Equal(t, nilSlice, val)

	// 2. []byte(nil) literal
	err = multipart.AddOrReplacePart("binaryPart2", "application/octet-stream", []byte(nil))
	assert.NoError(t, err)

	// 3. empty allocated []byte{}
	err = multipart.AddOrReplacePart("binaryPart3", "application/octet-stream", []byte{})
	assert.NoError(t, err)

	// 4. untyped nil should still be rejected
	err = multipart.AddOrReplacePart("invalidPart", "application/octet-stream", nil)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "content cannot be nil")

	// 5. typed nil pointer should still be rejected
	var nilParsable *dummyParsable
	err = multipart.AddOrReplacePart("invalidPart2", "application/json", nilParsable)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "content cannot be nil")
}


func TestMultipartBodySerializeWithEmptyBinaryPart(t *testing.T) {
	multipart := NewMultipartBody()
	mockAdapter := &MockRequestAdapter{
		SerializationWriterFactory: &errorMockSerializerFactory{},
	}
	multipart.SetRequestAdapter(mockAdapter)

	var emptyBytes []byte
	err := multipart.AddOrReplacePart("file", "application/octet-stream", emptyBytes)
	assert.NoError(t, err)

	mockWriter := &internal.MockSerializer{
		CallsCounter: make(map[string]int),
	}

	err = multipart.Serialize(mockWriter)
	assert.NoError(t, err)
	assert.Equal(t, 1, mockWriter.CallsCounter["WriteByteArrayValue"])
}
