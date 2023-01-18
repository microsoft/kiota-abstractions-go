package store

// InMemoryBackingStoreFactory is an implementation of BackingStoreFactory and initializes an instance of inMemoryBackingStore
type InMemoryBackingStoreFactory struct {
}

func (i *InMemoryBackingStoreFactory) CreateBackingStore() BackingStore {
	return NewInMemoryBackingStore()
}
