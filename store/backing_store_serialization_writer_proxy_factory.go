package store

import (
	"errors"
	"strings"

	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// BackingStoreSerializationWriterProxyFactory Backing Store implementation for serialization.SerializationWriterFactory
type BackingStoreSerializationWriterProxyFactory struct {
	factory serialization.SerializationWriterFactory
}

func (b *BackingStoreSerializationWriterProxyFactory) GetValidContentType() (string, error) {
	if b == nil || b.factory == nil {
		return "", errors.New("factory is nil")
	}
	return b.factory.GetValidContentType()
}

func (b *BackingStoreSerializationWriterProxyFactory) GetSerializationWriter(contentType string) (serialization.SerializationWriter, error) {
	if b == nil || b.factory == nil {
		return nil, errors.New("factory is nil")
	}
	return b.factory.GetSerializationWriter(contentType)
}

// NewBackingStoreSerializationWriterProxyFactory Initializes a new instance of BackingStoreSerializationWriterProxyFactory
func NewBackingStoreSerializationWriterProxyFactory(factory serialization.SerializationWriterFactory) *BackingStoreSerializationWriterProxyFactory {
	if b, ok := factory.(*BackingStoreSerializationWriterProxyFactory); ok {
		if b == nil {
			return nil
		}
		return b
	}
	if factory == nil {
		return nil
	}
	proxyFactory := serialization.NewSerializationWriterProxyFactory(factory, func(parsable serialization.Parsable) error {
		if backedModel, ok := parsable.(BackedModel); ok && backedModel.GetBackingStore() != nil {
			backedModel.GetBackingStore().SetReturnOnlyChangedValues(true)
		}
		return nil
	}, func(parsable serialization.Parsable) error {
		if backedModel, ok := parsable.(BackedModel); ok && backedModel.GetBackingStore() != nil {
			store := backedModel.GetBackingStore()
			store.SetReturnOnlyChangedValues(false)
			store.SetInitializationCompleted(true)
		}
		return nil
	}, func(parsable serialization.Parsable, writer serialization.SerializationWriter) error {
		if backedModel, ok := parsable.(BackedModel); ok && backedModel.GetBackingStore() != nil {

			nilValues := backedModel.GetBackingStore().EnumerateKeysForValuesChangedToNil()
			if len(nilValues) == 0 {
				return nil
			}
			var deserializers map[string]func(serialization.ParseNode) error
			if parsable != nil {
				deserializers = parsable.GetFieldDeserializers()
			}
			for _, k := range nilValues {
				resolvedKey := resolveSerializedKey(k, deserializers)
				err := writer.WriteNullValue(resolvedKey)
				if err != nil {
					return err
				}
			}
		}

		return nil
	})

	return &BackingStoreSerializationWriterProxyFactory{
		factory: proxyFactory,
	}
}

func normalizeKey(s string) string {
	var b strings.Builder
	b.Grow(len(s))
	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			b.WriteRune(r)
		} else if r >= 'A' && r <= 'Z' {
			b.WriteRune(r + ('a' - 'A'))
		}
	}
	return b.String()
}

func resolveSerializedKey(k string, deserializers map[string]func(serialization.ParseNode) error) string {
	if len(deserializers) == 0 {
		return k
	}
	if _, ok := deserializers[k]; ok {
		return k
	}
	for propName := range deserializers {
		if strings.EqualFold(propName, k) {
			return propName
		}
	}
	normK := normalizeKey(k)
	for propName := range deserializers {
		if normalizeKey(propName) == normK {
			return propName
		}
	}
	trimmedK := strings.TrimPrefix(normK, "escaped")
	trimmedK = strings.TrimSuffix(trimmedK, "escaped")
	if trimmedK != normK && trimmedK != "" {
		for propName := range deserializers {
			if normalizeKey(propName) == trimmedK {
				return propName
			}
		}
	}
	return k
}
