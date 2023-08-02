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
