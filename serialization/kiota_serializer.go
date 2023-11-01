package serialization

import (
	"errors"
	"reflect"
	"strings"
)

func Serialize(contentType string, model Parsable) ([]byte, error) {
	writer, err := getSerializationWriter(contentType, model)
	if err != nil {
		return nil, err
	}
	defer writer.Close()
	err = writer.WriteObjectValue("", model)
	if err != nil {
		return nil, err
	}
	return writer.GetSerializedContent()
}
func SerializeCollection(contentType string, models []Parsable) ([]byte, error) {
	writer, err := getSerializationWriter(contentType, models)
	if err != nil {
		return nil, err
	}
	defer writer.Close()
	err = writer.WriteCollectionOfObjectValues("", models)
	if err != nil {
		return nil, err
	}
	return writer.GetSerializedContent()
}
func getSerializationWriter(contentType string, value interface{}) (SerializationWriter, error) {
	if contentType == "" {
		return nil, errors.New("the content type is empty")
	}
	if value == nil {
		return nil, errors.New("the value is empty")
	}
	writer, err := DefaultSerializationWriterFactoryInstance.GetSerializationWriter(contentType)
	if err != nil {
		return nil, err
	}
	return writer, nil
}

func Deserialize(contentType string, content []byte, parsableFactory ParsableFactory) (Parsable, error) {
	node, err := getParseNode(contentType, content, parsableFactory)
	if err != nil {
		return nil, err
	}
	result, err := node.GetObjectValue(parsableFactory)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func getParseNode(contentType string, content []byte, parsableFactory ParsableFactory) (ParseNode, error) {
	if contentType == "" {
		return nil, errors.New("the content type is empty")
	}
	if content == nil || len(content) == 0 {
		return nil, errors.New("the content is empty")
	}
	if parsableFactory == nil {
		return nil, errors.New("the parsable factory is empty")
	}
	node, err := DefaultParseNodeFactoryInstance.GetRootParseNode(contentType, content)
	if err != nil {
		return nil, err
	}
	return node, nil
}
func DeserializeCollection(contentType string, content []byte, parsableFactory ParsableFactory) ([]Parsable, error) {
	node, err := getParseNode(contentType, content, parsableFactory)
	if err != nil {
		return nil, err
	}
	result, err := node.GetCollectionOfObjectValues(parsableFactory)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func DeserializeWithType(contentType string, content []byte, modelType reflect.Type) (Parsable, error) {
	factory, err := getTypeFactory(modelType)
	if err != nil {
		return nil, err
	}
	return Deserialize(contentType, content, factory)
}
func DeserializeCollectionWithType(contentType string, content []byte, modelType reflect.Type) ([]Parsable, error) {
	factory, err := getTypeFactory(modelType)
	if err != nil {
		return nil, err
	}
	return DeserializeCollection(contentType, content, factory)
}

func getTypeFactory(modelType reflect.Type) (ParsableFactory, error) {
	typeName := modelType.Name()
	method, ok := modelType.MethodByName("Create" + strings.ToUpper(typeName[0:0]) + typeName[1:] + "FromDiscriminatorValue")
	if !ok {
		return nil, errors.New("the model type does not have a factory method")
	}
	factory := func(parseNode ParseNode) (Parsable, error) {
		return method.Func.Call([]reflect.Value{reflect.ValueOf(parseNode)})[0].Interface().(Parsable), nil
	}
	return factory, nil
}
