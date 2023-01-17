package store

import "github.com/microsoft/kiota-abstractions-go/serialization"

func NewBackingStoreParseNodeFactory(factory serialization.ParseNodeFactory) *serialization.ParseNodeProxyFactory {
	return serialization.NewParseNodeProxyFactory(factory, func(parsable serialization.Parsable) {
		if backedModel, ok := parsable.(BackedModel); ok && backedModel.GetBackingStore() != nil {
			initialization := false
			(*backedModel.GetBackingStore()).SetInitializationCompleted(&initialization)
		}
	}, func(parsable serialization.Parsable) {
		if backedModel, ok := parsable.(BackedModel); ok && backedModel.GetBackingStore() != nil {
			initialization := true
			(*backedModel.GetBackingStore()).SetInitializationCompleted(&initialization)
		}
	})
}
