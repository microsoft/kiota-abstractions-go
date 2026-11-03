package store

import (
	"errors"

	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// BackingStoreParseNodeFactory Backing Store implementation for serialization.ParseNodeFactory
type BackingStoreParseNodeFactory struct {
	serialization.ParseNodeFactory
}

func (b *BackingStoreParseNodeFactory) GetValidContentType() (string, error) {
	if b == nil || b.ParseNodeFactory == nil {
		return "", errors.New("factory is nil")
	}
	return b.ParseNodeFactory.GetValidContentType()
}

func (b *BackingStoreParseNodeFactory) GetRootParseNode(contentType string, content []byte) (serialization.ParseNode, error) {
	if b == nil || b.ParseNodeFactory == nil {
		return nil, errors.New("factory is nil")
	}
	return b.ParseNodeFactory.GetRootParseNode(contentType, content)
}

// NewBackingStoreParseNodeFactory Initializes a new instance of BackingStoreParseNodeFactory
func NewBackingStoreParseNodeFactory(factory serialization.ParseNodeFactory) *BackingStoreParseNodeFactory {
	proxyFactory := serialization.NewParseNodeProxyFactory(factory, func(parsable serialization.Parsable) error {
		if backedModel, ok := parsable.(BackedModel); ok && backedModel.GetBackingStore() != nil {
			backedModel.GetBackingStore().SetInitializationCompleted(false)
		}
		return nil
	}, func(parsable serialization.Parsable) error {
		if backedModel, ok := parsable.(BackedModel); ok && backedModel.GetBackingStore() != nil {
			backedModel.GetBackingStore().SetInitializationCompleted(true)
		}
		return nil
	})

	return &BackingStoreParseNodeFactory{proxyFactory}
}
