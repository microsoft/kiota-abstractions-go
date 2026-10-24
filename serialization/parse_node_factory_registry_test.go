package serialization

import (
	"testing"

	assert "github.com/stretchr/testify/assert"
)

func TestParseNodeFactoryRegistryHonoursInterface(t *testing.T) {
	assert.Implements(t, (*ParseNodeFactory)(nil), DefaultParseNodeFactoryInstance)
}

type stubParseNodeFactory struct {
	contentType string
}

func (s *stubParseNodeFactory) GetValidContentType() (string, error) {
	return s.contentType, nil
}

func (s *stubParseNodeFactory) GetRootParseNode(contentType string, content []byte) (ParseNode, error) {
	return nil, nil
}

type stubWriterFactory struct {
	contentType string
}

func (s *stubWriterFactory) GetValidContentType() (string, error) {
	return s.contentType, nil
}

func (s *stubWriterFactory) GetSerializationWriter(contentType string) (SerializationWriter, error) {
	return nil, nil
}

func TestParseNodeFactoryRegistryNilGuards(t *testing.T) {
	var nilRegistry *ParseNodeFactoryRegistry
	assert.NotPanics(t, func() {
		node, err := nilRegistry.GetRootParseNode("application/json", []byte("{}"))
		assert.Error(t, err)
		assert.Nil(t, node)
		nilRegistry.Lock()
		nilRegistry.Unlock()
	})
}

func TestSerializationWriterFactoryRegistryNilGuards(t *testing.T) {
	var nilRegistry *SerializationWriterFactoryRegistry
	assert.NotPanics(t, func() {
		writer, err := nilRegistry.GetSerializationWriter("application/json")
		assert.Error(t, err)
		assert.Nil(t, writer)
		nilRegistry.Lock()
		nilRegistry.Unlock()
	})
}

func TestFactoryRegistryContentTypeTolerance(t *testing.T) {
	parseRegistry := NewParseNodeFactoryRegistry()
	stubNode := &stubParseNodeFactory{contentType: "application/json"}
	parseRegistry.ContentTypeAssociatedFactories["application/json"] = stubNode

	// Case-insensitivity & whitespace trimming
	_, err := parseRegistry.GetRootParseNode("  APPLICATION/JSON; charset=utf-8  ", []byte("{}"))
	assert.NoError(t, err)

	_, errEmpty := parseRegistry.GetRootParseNode("   ", []byte("{}"))
	assert.Error(t, errEmpty)

	writerRegistry := NewSerializationWriterFactoryRegistry()
	stubWriter := &stubWriterFactory{contentType: "application/json"}
	writerRegistry.ContentTypeAssociatedFactories["application/json"] = stubWriter

	_, errW := writerRegistry.GetSerializationWriter("  APPLICATION/JSON; charset=utf-8  ")
	assert.NoError(t, errW)

	_, errWEmpty := writerRegistry.GetSerializationWriter("   ")
	assert.Error(t, errWEmpty)
}
