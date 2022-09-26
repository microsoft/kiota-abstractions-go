package serialization

import (
	"github.com/microsoft/kiota-abstractions-go/internal"
	assert "github.com/stretchr/testify/assert"
	"testing"
)

func TestItGetsVendorSpecificSerializationWriter(t *testing.T) {
	DefaultSerializationWriterFactoryInstance.ContentTypeAssociatedFactories["application/json"] = &internal.MockSerializerFactory{}
	serializationWriter, err := DefaultSerializationWriterFactoryInstance.GetSerializationWriter("application/vnd+json")
	assert.Nil(t, err)
	assert.NotNil(t, serializationWriter)
}

func TestSerializationWriterFactoryRegistryHonoursInterface(t *testing.T) {
	assert.Implements(t, (*SerializationWriterFactory)(nil), DefaultSerializationWriterFactoryInstance)
}
