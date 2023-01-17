package internal

import (
	"time"

	"github.com/google/uuid"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

type MockSerializer struct {
	CallsCounter map[string]int
}

func (m *MockSerializer) WriteNullValue(key string) error {
	return nil
}

func (m *MockSerializer) GetOnBeforeSerialization() serialization.ParsableAction {
	return nil
}

func (m *MockSerializer) SetOnBeforeSerialization(action serialization.ParsableAction) {

}

func (m *MockSerializer) GetOnAfterObjectSerialization() serialization.ParsableAction {
	return nil
}

func (m *MockSerializer) SetOnAfterObjectSerialization(action serialization.ParsableAction) {

}

func (m *MockSerializer) GetOnStartObjectSerialization() serialization.ParsableWriter {
	return nil
}

func (m *MockSerializer) SetOnStartObjectSerialization(writer serialization.ParsableWriter) {

}

func (m *MockSerializer) WriteStringValue(key string, value *string) error {
	m.CallsCounter["WriteStringValue"]++
	return nil
}
func (*MockSerializer) WriteBoolValue(key string, value *bool) error {
	return nil
}
func (*MockSerializer) WriteByteValue(key string, value *byte) error {
	return nil
}
func (*MockSerializer) WriteInt8Value(key string, value *int8) error {
	return nil
}
func (*MockSerializer) WriteInt32Value(key string, value *int32) error {
	return nil
}
func (*MockSerializer) WriteInt64Value(key string, value *int64) error {
	return nil
}
func (*MockSerializer) WriteFloat32Value(key string, value *float32) error {
	return nil
}
func (*MockSerializer) WriteFloat64Value(key string, value *float64) error {
	return nil
}
func (*MockSerializer) WriteByteArrayValue(key string, value []byte) error {
	return nil
}
func (*MockSerializer) WriteTimeValue(key string, value *time.Time) error {
	return nil
}
func (*MockSerializer) WriteISODurationValue(key string, value *serialization.ISODuration) error {
	return nil
}
func (*MockSerializer) WriteDateOnlyValue(key string, value *serialization.DateOnly) error {
	return nil
}
func (*MockSerializer) WriteTimeOnlyValue(key string, value *serialization.TimeOnly) error {
	return nil
}
func (*MockSerializer) WriteUUIDValue(key string, value *uuid.UUID) error {
	return nil
}
func (m *MockSerializer) WriteObjectValue(key string, item serialization.Parsable, additionalValuesToMerge ...serialization.Parsable) error {
	m.CallsCounter["WriteObjectValue"]++
	return nil
}
func (m *MockSerializer) WriteCollectionOfObjectValues(key string, collection []serialization.Parsable) error {
	m.CallsCounter["WriteCollectionOfObjectValues"]++
	return nil
}
func (m *MockSerializer) WriteCollectionOfStringValues(key string, collection []string) error {
	m.CallsCounter["WriteCollectionOfStringValues"]++
	return nil
}
func (*MockSerializer) WriteCollectionOfBoolValues(key string, collection []bool) error {
	return nil
}
func (*MockSerializer) WriteCollectionOfByteValues(key string, collection []byte) error {
	return nil
}
func (*MockSerializer) WriteCollectionOfInt8Values(key string, collection []int8) error {
	return nil
}
func (*MockSerializer) WriteCollectionOfInt32Values(key string, collection []int32) error {
	return nil
}
func (*MockSerializer) WriteCollectionOfInt64Values(key string, collection []int64) error {
	return nil
}
func (*MockSerializer) WriteCollectionOfFloat32Values(key string, collection []float32) error {
	return nil
}
func (*MockSerializer) WriteCollectionOfFloat64Values(key string, collection []float64) error {
	return nil
}
func (*MockSerializer) WriteCollectionOfTimeValues(key string, collection []time.Time) error {
	return nil
}
func (*MockSerializer) WriteCollectionOfISODurationValues(key string, collection []serialization.ISODuration) error {
	return nil
}
func (*MockSerializer) WriteCollectionOfDateOnlyValues(key string, collection []serialization.DateOnly) error {
	return nil
}
func (*MockSerializer) WriteCollectionOfTimeOnlyValues(key string, collection []serialization.TimeOnly) error {
	return nil
}
func (*MockSerializer) WriteCollectionOfUUIDValues(key string, collection []uuid.UUID) error {
	return nil
}
func (*MockSerializer) GetSerializedContent() ([]byte, error) {
	return []byte("content"), nil
}
func (*MockSerializer) WriteAdditionalData(value map[string]interface{}) error {
	return nil
}
func (*MockSerializer) WriteAnyValue(key string, value interface{}) error {
	return nil
}
func (*MockSerializer) Close() error {
	return nil
}

type MockSerializerFactory struct {
	SerializationWriter serialization.SerializationWriter
}

func (*MockSerializerFactory) GetValidContentType() (string, error) {
	return "application/json", nil
}
func (m *MockSerializerFactory) GetSerializationWriter(contentType string) (serialization.SerializationWriter, error) {
	if m.SerializationWriter == nil {
		m.SerializationWriter = &MockSerializer{
			CallsCounter: make(map[string]int),
		}
	}
	return m.SerializationWriter, nil
}
