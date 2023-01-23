package abstractions

import (
	"fmt"
	"github.com/microsoft/kiota-abstractions-go/store"
	"reflect"
	"testing"
	"time"

	"github.com/microsoft/kiota-abstractions-go/internal"
	serialization "github.com/microsoft/kiota-abstractions-go/serialization"
	assert "github.com/stretchr/testify/assert"
)

func TestItCreatesClientConcurrently(t *testing.T) {
	metaFactory := func() serialization.SerializationWriterFactory {
		return &internal.MockSerializerFactory{}
	}
	for i := 0; i < 1000; i++ {
		go RegisterDefaultSerializer(metaFactory)
	}
	waitTime, _ := time.ParseDuration("100ms") // otherwise the routines might not be completed
	time.Sleep(waitTime)
	assert.Equal(t, 1, len(serialization.DefaultSerializationWriterFactoryInstance.ContentTypeAssociatedFactories))
}

func TestEnableBackingStoreForSerializationWriterFactory(t *testing.T) {
	serializationFactoryRegistry := serialization.DefaultSerializationWriterFactoryInstance
	factory := &internal.MockSerializerFactory{}
	StreamContentType := "application/octet-stream"

	serializationFactoryRegistry.ContentTypeAssociatedFactories[StreamContentType] = factory

	IsNotType(t, &store.BackingStoreSerializationWriterProxyFactory{}, serializationFactoryRegistry.ContentTypeAssociatedFactories[StreamContentType])

	EnableBackingStoreForSerializationWriterFactory(factory)
	assert.IsType(t, &store.BackingStoreSerializationWriterProxyFactory{}, serializationFactoryRegistry.ContentTypeAssociatedFactories[StreamContentType])
}

func TestEnableBackingStoreForParseNodeFactory(t *testing.T) {
	parseNodeRegistry := serialization.DefaultParseNodeFactoryInstance
	factory := internal.NewMockParseNodeFactory()
	StreamContentType := "application/octet-stream"

	parseNodeRegistry.ContentTypeAssociatedFactories[StreamContentType] = factory

	IsNotType(t, &store.BackingStoreParseNodeFactory{}, parseNodeRegistry.ContentTypeAssociatedFactories[StreamContentType])

	EnableBackingStoreForParseNodeFactory(factory)
	assert.IsType(t, &store.BackingStoreParseNodeFactory{}, parseNodeRegistry.ContentTypeAssociatedFactories[StreamContentType])
}

// IsType asserts that the specified objects are of the same type.
func IsNotType(t assert.TestingT, expectedType interface{}, object interface{}, msgAndArgs ...interface{}) bool {
	if assert.ObjectsAreEqual(reflect.TypeOf(object), reflect.TypeOf(expectedType)) {
		return assert.Fail(t, fmt.Sprintf("Object expected to be of type %v, but was %v", reflect.TypeOf(expectedType), reflect.TypeOf(object)), msgAndArgs...)
	}

	return true
}
