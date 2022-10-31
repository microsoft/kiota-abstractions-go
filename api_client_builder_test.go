package abstractions

import (
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
