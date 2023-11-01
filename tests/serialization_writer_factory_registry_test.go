package tests

import (
	"testing"

	"github.com/microsoft/kiota-abstractions-go/serialization"

	"github.com/microsoft/kiota-abstractions-go/internal"
	assert "github.com/stretchr/testify/assert"
)

func TestItGetsVendorSpecificSerializationWriter(t *testing.T) {
	serialization.DefaultSerializationWriterFactoryInstance.ContentTypeAssociatedFactories["application/json"] = &internal.MockSerializerFactory{}
	serializationWriter, err := serialization.DefaultSerializationWriterFactoryInstance.GetSerializationWriter("application/vnd+json")
	assert.Nil(t, err)
	assert.NotNil(t, serializationWriter)
	serialization.DefaultSerializationWriterFactoryInstance.ContentTypeAssociatedFactories = make(map[string]serialization.SerializationWriterFactory)
}

func TestSerializationWriterFactoryRegistryHonoursInterface(t *testing.T) {
	assert.Implements(t, (*serialization.SerializationWriterFactory)(nil), serialization.DefaultSerializationWriterFactoryInstance)
}
