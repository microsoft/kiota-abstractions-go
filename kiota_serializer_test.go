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

func TestItDefendsDeserializationEmptyContentType(t *testing.T) {
	result, err := serialization.Deserialize("", nil, nil)
	assert.Nil(t, result)
	assert.NotNil(t, err)
}
func TestItDefendsDeserializationNilContent(t *testing.T) {
	result, err := serialization.Deserialize(jsonContentType, nil, nil)
	assert.Nil(t, result)
	assert.NotNil(t, err)
}
func TestItDefendsDeserializationNilFactory(t *testing.T) {
	result, err := serialization.Deserialize(jsonContentType, make([]byte, 0), nil)
	assert.Nil(t, result)
	assert.NotNil(t, err)
}

func TestItDefendsCollectionDeserializationEmptyContentType(t *testing.T) {
	result, err := serialization.DeserializeCollection("", nil, nil)
	assert.Nil(t, result)
	assert.NotNil(t, err)
}
func TestItDefendsCollectionDeserializationNilContent(t *testing.T) {
	result, err := serialization.DeserializeCollection(jsonContentType, nil, nil)
	assert.Nil(t, result)
	assert.NotNil(t, err)
}
func TestItDefendsCollectionDeserializationNilFactory(t *testing.T) {
	result, err := serialization.DeserializeCollection(jsonContentType, make([]byte, 0), nil)
	assert.Nil(t, result)
	assert.NotNil(t, err)
}

func TestItDefendsDeserializationNilType(t *testing.T) {
	result, err := serialization.DeserializeWithType(jsonContentType, make([]byte, 0), nil)
	assert.Nil(t, result)
	assert.NotNil(t, err)
}

func TestItDefendsCollectionDeserializationNilType(t *testing.T) {
	result, err := serialization.DeserializeCollectionWithType(jsonContentType, make([]byte, 0), nil)
	assert.Nil(t, result)
	assert.NotNil(t, err)
}

func TestItDeserializesAnObject(t *testing.T) {
	person := internal.NewPerson()
	id := "123"
	person.SetId(&id)
	metaFactory := func() serialization.ParseNodeFactory {
		return &internal.MockParseNodeFactory{
			SerializedValue: person,
		}
	}
	RegisterDefaultDeserializer(metaFactory)

	result, err := serialization.Deserialize(jsonContentType, []byte("{\"id\": \"123\"}"), internal.CreatePersonFromDiscriminatorValue)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	resultAsPerson, ok := result.(*internal.Person)
	assert.True(t, ok)
	assert.Equal(t, id, *resultAsPerson.GetId())
	serialization.DefaultParseNodeFactoryInstance.ContentTypeAssociatedFactories = make(map[string]serialization.ParseNodeFactory)
}

func TestItDeserializesAnObjectCollection(t *testing.T) {
	person := internal.NewPerson()
	id := "123"
	person.SetId(&id)
	metaFactory := func() serialization.ParseNodeFactory {
		return &internal.MockParseNodeFactory{
			SerializedValue: []serialization.Parsable{person},
		}
	}
	RegisterDefaultDeserializer(metaFactory)

	result, err := serialization.DeserializeCollection(jsonContentType, []byte("[{\"id\": \"123\"}]"), internal.CreatePersonFromDiscriminatorValue)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	resultAsPerson, ok := result[0].(*internal.Person)
	assert.True(t, ok)
	assert.Equal(t, id, *resultAsPerson.GetId())
	serialization.DefaultParseNodeFactoryInstance.ContentTypeAssociatedFactories = make(map[string]serialization.ParseNodeFactory)
}
