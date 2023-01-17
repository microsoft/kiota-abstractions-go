package store

import (
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

func NewBackingStoreSerializationWriterProxyFactory(factory serialization.SerializationWriterFactory) *serialization.SerializationWriterProxyFactory {
	return serialization.NewSerializationWriterProxyFactory(factory, func(parsable serialization.Parsable) {
		if backedModel, ok := parsable.(BackedModel); ok && backedModel.GetBackingStore() != nil {
			returnChanged := true
			(*backedModel.GetBackingStore()).SetReturnOnlyChangedValues(&returnChanged)
		}
	}, func(parsable serialization.Parsable) {
		if backedModel, ok := parsable.(BackedModel); ok && backedModel.GetBackingStore() != nil {
			store := *backedModel.GetBackingStore()
			returnChanged := false
			store.SetReturnOnlyChangedValues(&returnChanged)
			initializationComplete := true
			store.SetInitializationCompleted(&initializationComplete)
		}
	}, func(parsable serialization.Parsable, writer serialization.SerializationWriter) error {
		if backedModel, ok := parsable.(BackedModel); ok && backedModel.GetBackingStore() != nil {

			nilValues := (*backedModel.GetBackingStore()).EnumerateKeysForValuesChangedToNil()
			for _, k := range nilValues {
				err := writer.WriteNullValue(k)
				if err != nil {
					return err
				}
			}
		}

		return nil
	})
}
