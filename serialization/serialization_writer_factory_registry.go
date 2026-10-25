package serialization

import (
	"errors"
	"strings"
	"sync"
)

// SerializationWriterFactoryRegistry is a factory holds a list of all the registered factories for the various types of nodes.
type SerializationWriterFactoryRegistry struct {
	lock *sync.Mutex

	// ContentTypeAssociatedFactories list of factories that are registered by content type.
	//
	// When interacting with this field, please make use of Lock and Unlock methods to ensure thread safety.
	ContentTypeAssociatedFactories map[string]SerializationWriterFactory
}

func NewSerializationWriterFactoryRegistry() *SerializationWriterFactoryRegistry {
	return &SerializationWriterFactoryRegistry{
		lock:                           &sync.Mutex{},
		ContentTypeAssociatedFactories: make(map[string]SerializationWriterFactory),
	}
}

// DefaultSerializationWriterFactoryInstance is the default singleton instance of the registry to be used when registering new factories that should be available by default.
var DefaultSerializationWriterFactoryInstance = NewSerializationWriterFactoryRegistry()

// GetValidContentType returns the valid content type for the SerializationWriterFactoryRegistry
func (m *SerializationWriterFactoryRegistry) GetValidContentType() (string, error) {
	return "", errors.New("the registry supports multiple content types. Get the registered factory instead")
}

// GetSerializationWriter returns the relevant SerializationWriter instance for the given content type
func (m *SerializationWriterFactoryRegistry) GetSerializationWriter(contentType string) (SerializationWriter, error) {
	if m == nil {
		return nil, errors.New("registry is nil")
	}
	trimmed := strings.TrimSpace(contentType)
	if trimmed == "" {
		return nil, errors.New("the content type is empty")
	}
	vendorSpecificContentType := strings.TrimSpace(strings.Split(trimmed, ";")[0])
	factory, ok := m.ContentTypeAssociatedFactories[vendorSpecificContentType]
	if ok {
		return factory.GetSerializationWriter(normalizeWithParameters(vendorSpecificContentType, trimmed))
	}
	if matchedKey, factory, ok := m.getFactoryCaseInsensitive(vendorSpecificContentType); ok {
		return factory.GetSerializationWriter(normalizeWithParameters(matchedKey, trimmed))
	}
	cleanedContentType := contentTypeVendorCleanupPattern.ReplaceAllString(vendorSpecificContentType, "")
	factory, ok = m.ContentTypeAssociatedFactories[cleanedContentType]
	if ok {
		return factory.GetSerializationWriter(cleanedContentType)
	}
	if matchedKey, factory, ok := m.getFactoryCaseInsensitive(cleanedContentType); ok {
		return factory.GetSerializationWriter(matchedKey)
	}
	return nil, errors.New("Content type " + cleanedContentType + " does not have a factory registered to be parsed")
}

func normalizeWithParameters(matchedKey string, trimmed string) string {
	parts := strings.SplitN(trimmed, ";", 2)
	if len(parts) > 1 {
		return matchedKey + ";" + parts[1]
	}
	return matchedKey
}

func (m *SerializationWriterFactoryRegistry) getFactoryCaseInsensitive(target string) (string, SerializationWriterFactory, bool) {
	lowerTarget := strings.ToLower(target)
	for k, v := range m.ContentTypeAssociatedFactories {
		if strings.ToLower(k) == lowerTarget {
			return k, v, true
		}
	}
	return "", nil, false
}

func (m *SerializationWriterFactoryRegistry) Lock() {
	if m == nil || m.lock == nil {
		return
	}
	m.lock.Lock()
}

func (m *SerializationWriterFactoryRegistry) Unlock() {
	if m == nil || m.lock == nil {
		return
	}
	m.lock.Unlock()
}
