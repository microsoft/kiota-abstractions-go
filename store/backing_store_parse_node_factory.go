package store

import "github.com/microsoft/kiota-abstractions-go/serialization"

// BackingStoreParseNodeFactory Backing Store implementation for serialization.ParseNodeFactory
type BackingStoreParseNodeFactory struct {
	factory *serialization.ParseNodeProxyFactory
}

func (b *BackingStoreParseNodeFactory) GetValidContentType() (string, error) {
	return b.factory.GetValidContentType()
}

func (b *BackingStoreParseNodeFactory) GetRootParseNode(contentType string, content []byte) (serialization.ParseNode, error) {
	return b.factory.GetRootParseNode(contentType, content)
}

// NewBackingStoreParseNodeFactory Initializes a new instance of BackingStoreParseNodeFactory
func NewBackingStoreParseNodeFactory(factory serialization.ParseNodeFactory) *BackingStoreParseNodeFactory {
	proxyFactory := serialization.NewParseNodeProxyFactory(factory, func(parsable serialization.Parsable) {
		if backedModel, ok := parsable.(BackedModel); ok && backedModel.GetBackingStore() != nil {
			backedModel.GetBackingStore().SetInitializationCompleted(false)
		}
	}, func(parsable serialization.Parsable) {
		if backedModel, ok := parsable.(BackedModel); ok && backedModel.GetBackingStore() != nil {
			backedModel.GetBackingStore().SetInitializationCompleted(true)
		}
	})

	return &BackingStoreParseNodeFactory{
		factory: proxyFactory,
	}
}