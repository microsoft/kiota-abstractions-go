package store

type InMemoryBackingStoreFactory struct {
}

func (i *InMemoryBackingStoreFactory) CreateBackingStore() BackingStore {
	return NewInMemoryBackingStore()
}
