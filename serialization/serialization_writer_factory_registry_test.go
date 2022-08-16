package serialization

import (
	"testing"
	"time"

	"github.com/google/uuid"

	assert "github.com/stretchr/testify/assert"
)

type mockSerializer struct {
}

func (*mockSerializer) WriteStringValue(key string, value *string) error {
	return nil
}
func (*mockSerializer) WriteBoolValue(key string, value *bool) error {
	return nil
}
func (*mockSerializer) WriteByteValue(key string, value *byte) error {
	return nil
}
func (*mockSerializer) WriteInt8Value(key string, value *int8) error {
	return nil
}
func (*mockSerializer) WriteInt32Value(key string, value *int32) error {
	return nil
}
func (*mockSerializer) WriteInt64Value(key string, value *int64) error {
	return nil
}
func (*mockSerializer) WriteFloat32Value(key string, value *float32) error {
	return nil
}
func (*mockSerializer) WriteFloat64Value(key string, value *float64) error {
	return nil
}
func (*mockSerializer) WriteByteArrayValue(key string, value []byte) error {
	return nil
}
func (*mockSerializer) WriteTimeValue(key string, value *time.Time) error {
	return nil
}
func (*mockSerializer) WriteISODurationValue(key string, value *ISODuration) error {
	return nil
}
func (*mockSerializer) WriteDateOnlyValue(key string, value *DateOnly) error {
	return nil
}
func (*mockSerializer) WriteTimeOnlyValue(key string, value *TimeOnly) error {
	return nil
}
func (*mockSerializer) WriteUUIDValue(key string, value *uuid.UUID) error {
	return nil
}
func (*mockSerializer) WriteObjectValue(key string, item Parsable, additionalValuesToMerge ...Parsable) error {
	return nil
}
func (*mockSerializer) WriteCollectionOfObjectValues(key string, collection []Parsable) error {
	return nil
}
func (*mockSerializer) WriteCollectionOfStringValues(key string, collection []string) error {
	return nil
}
func (*mockSerializer) WriteCollectionOfBoolValues(key string, collection []bool) error {
	return nil
}
func (*mockSerializer) WriteCollectionOfByteValues(key string, collection []byte) error {
	return nil
}
func (*mockSerializer) WriteCollectionOfInt8Values(key string, collection []int8) error {
	return nil
}
func (*mockSerializer) WriteCollectionOfInt32Values(key string, collection []int32) error {
	return nil
}
func (*mockSerializer) WriteCollectionOfInt64Values(key string, collection []int64) error {
	return nil
}
func (*mockSerializer) WriteCollectionOfFloat32Values(key string, collection []float32) error {
	return nil
}
func (*mockSerializer) WriteCollectionOfFloat64Values(key string, collection []float64) error {
	return nil
}
func (*mockSerializer) WriteCollectionOfTimeValues(key string, collection []time.Time) error {
	return nil
}
func (*mockSerializer) WriteCollectionOfISODurationValues(key string, collection []ISODuration) error {
	return nil
}
func (*mockSerializer) WriteCollectionOfDateOnlyValues(key string, collection []DateOnly) error {
	return nil
}
func (*mockSerializer) WriteCollectionOfTimeOnlyValues(key string, collection []TimeOnly) error {
	return nil
}
func (*mockSerializer) WriteCollectionOfUUIDValues(key string, collection []uuid.UUID) error {
	return nil
}
func (*mockSerializer) GetSerializedContent() ([]byte, error) {
	return nil, nil
}
func (*mockSerializer) WriteAdditionalData(value map[string]interface{}) error {
	return nil
}
func (*mockSerializer) Close() error {
	return nil
}

type mockSerializerFactory struct {
}

func (*mockSerializerFactory) GetValidContentType() (string, error) {
	return "application/json", nil
}
func (*mockSerializerFactory) GetSerializationWriter(contentType string) (SerializationWriter, error) {
	return &mockSerializer{}, nil
}

func TestItGetsVendorSpecificSerializationWriter(t *testing.T) {
	DefaultSerializationWriterFactoryInstance.ContentTypeAssociatedFactories["application/json"] = &mockSerializerFactory{}
	serializationWriter, err := DefaultSerializationWriterFactoryInstance.GetSerializationWriter("application/vnd+json")
	assert.Nil(t, err)
	assert.NotNil(t, serializationWriter)
}

func TestSerializationWriterFactoryRegistryHonoursInterface(t *testing.T) {
	assert.Implements(t, (*SerializationWriterFactory)(nil), DefaultSerializationWriterFactoryInstance)
}
