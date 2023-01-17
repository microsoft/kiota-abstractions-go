package store

type BackingStoreFactory interface {
	// CreateBackingStore initializes a new backing store
	CreateBackingStore() BackingStore
}
