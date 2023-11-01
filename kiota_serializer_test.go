package abstractions

import (
	"testing"

	"github.com/microsoft/kiota-abstractions-go/internal"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	assert "github.com/stretchr/testify/assert"
)

var jsonContentType = "application/json"

func TestItDefendsSerializationEmptyContentType(t *testing.T) {
	result, err := serialization.Serialize("", nil)
	assert.Nil(t, result)
	assert.NotNil(t, err)
}

func TestItDefendsSerializationNilValue(t *testing.T) {
	result, err := serialization.Serialize(jsonContentType, nil)
	assert.Nil(t, result)
	assert.NotNil(t, err)
}

func TestItDefendsCollectionSerializationEmptyContentType(t *testing.T) {
	result, err := serialization.SerializeCollection("", nil)
	assert.Nil(t, result)
	assert.NotNil(t, err)
}

func TestItDefendsCollectionSerializationNilValue(t *testing.T) {
	result, err := serialization.SerializeCollection(jsonContentType, nil)
	assert.Nil(t, result)
	assert.NotNil(t, err)
}

func TestItSerializesObject(t *testing.T) {
	serializedValue := "{\"id\":\"123\"}"
	metaFactory := func() serialization.SerializationWriterFactory {
		return &internal.MockSerializerFactory{
			SerializedValue: serializedValue,
		}
	}
	RegisterDefaultSerializer(metaFactory)
	person := internal.NewPerson()
	id := "123"
	person.SetId(&id)
	result, err := serialization.Serialize(jsonContentType, person)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, serializedValue, string(result))
	serialization.DefaultSerializationWriterFactoryInstance.ContentTypeAssociatedFactories = make(map[string]serialization.SerializationWriterFactory)
}

func TestItSerializesACollectionOfObjects(t *testing.T) {
	serializedValue := "[{\"id\":\"123\"}]"
	metaFactory := func() serialization.SerializationWriterFactory {
		return &internal.MockSerializerFactory{
			SerializedValue: serializedValue,
		}
	}
	RegisterDefaultSerializer(metaFactory)
	person := internal.NewPerson()
	id := "123"
	person.SetId(&id)
	result, err := serialization.SerializeCollection(jsonContentType, []serialization.Parsable{person})
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, serializedValue, string(result))
	serialization.DefaultSerializationWriterFactoryInstance.ContentTypeAssociatedFactories = make(map[string]serialization.SerializationWriterFactory)
}
