package serialization

import (
	"errors"
	"testing"

	assert "github.com/stretchr/testify/assert"
)

func TestParseNodeFactoryRegistryHonoursInterface(t *testing.T) {
	assert.Implements(t, (*ParseNodeFactory)(nil), DefaultParseNodeFactoryInstance)
}

type validatingParseNodeFactory struct {
	expectedContentType string
	calledWith          string
}

func (s *validatingParseNodeFactory) GetValidContentType() (string, error) {
	return s.expectedContentType, nil
}

func (s *validatingParseNodeFactory) GetRootParseNode(contentType string, content []byte) (ParseNode, error) {
	s.calledWith = contentType
	if contentType != s.expectedContentType {
		return nil, errors.New("expected contentType: " + s.expectedContentType + ", got: " + contentType)
	}
	return nil, nil
}

type validatingWriterFactory struct {
	expectedContentType string
	calledWith          string
}

func (s *validatingWriterFactory) GetValidContentType() (string, error) {
	return s.expectedContentType, nil
}

func (s *validatingWriterFactory) GetSerializationWriter(contentType string) (SerializationWriter, error) {
	s.calledWith = contentType
	if contentType != s.expectedContentType {
		return nil, errors.New("expected contentType: " + s.expectedContentType + ", got: " + contentType)
	}
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
	stubNode := &validatingParseNodeFactory{expectedContentType: "application/json"}
	parseRegistry.ContentTypeAssociatedFactories["application/json"] = stubNode

	// Whitespace trimming & parameters
	_, err := parseRegistry.GetRootParseNode("  application/json; charset=utf-8  ", []byte("{}"))
	assert.NoError(t, err)
	assert.Equal(t, "application/json", stubNode.calledWith)

	// Mixed-case resolution
	_, errUpper := parseRegistry.GetRootParseNode("  APPLICATION/JSON; charset=utf-8  ", []byte("{}"))
	assert.NoError(t, errUpper)
	assert.Equal(t, "application/json", stubNode.calledWith)

	// Vendor cleanup fallback with mixed-case
	_, errVendor := parseRegistry.GetRootParseNode("  APPLICATION/VND.GITHUB+JSON; charset=utf-8  ", []byte("{}"))
	assert.NoError(t, errVendor)
	assert.Equal(t, "application/json", stubNode.calledWith)

	_, errEmpty := parseRegistry.GetRootParseNode("   ", []byte("{}"))
	assert.Error(t, errEmpty)

	writerRegistry := NewSerializationWriterFactoryRegistry()
	stubWriter := &validatingWriterFactory{expectedContentType: "application/json"}
	writerRegistry.ContentTypeAssociatedFactories["application/json"] = stubWriter

	// Direct match with whitespace and parameters
	_, errW := writerRegistry.GetSerializationWriter("  application/json; charset=utf-8  ")
	assert.NoError(t, errW)
	assert.Equal(t, "application/json", stubWriter.calledWith)

	// Direct match with mixed-case and parameters
	_, errWUpper := writerRegistry.GetSerializationWriter("  APPLICATION/JSON; charset=utf-8  ")
	assert.NoError(t, errWUpper)
	assert.Equal(t, "application/json", stubWriter.calledWith)

	// Direct match without parameters
	stubWriter.expectedContentType = "application/json"
	_, errWNoParams := writerRegistry.GetSerializationWriter("application/json")
	assert.NoError(t, errWNoParams)
	assert.Equal(t, "application/json", stubWriter.calledWith)

	// Vendor cleanup fallback with mixed-case
	_, errWVendor := writerRegistry.GetSerializationWriter("  APPLICATION/VND.GITHUB+JSON; charset=utf-8  ")
	assert.NoError(t, errWVendor)
	assert.Equal(t, "application/json", stubWriter.calledWith)

	_, errWEmpty := writerRegistry.GetSerializationWriter("   ")
	assert.Error(t, errWEmpty)
}
