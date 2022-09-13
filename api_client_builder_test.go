package abstractions

import (
	"testing"

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
	assert.Equal(t, 1, len(serialization.DefaultSerializationWriterFactoryInstance.ContentTypeAssociatedFactories))
}
