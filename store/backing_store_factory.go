package store

import "sync"

type BackingStoreFactory interface {
	// CreateBackingStore initializes a new backing store
	CreateBackingStore() BackingStore
}

var lock = &sync.Mutex{}

var singleInstance BackingStoreFactory

// GetDefaultBackingStoreInstance returns a backing store instance.
// if none exists an instance of inMemoryBackingStore is initialized and returned
func GetDefaultBackingStoreInstance() BackingStoreFactory {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		singleInstance = &InMemoryBackingStoreFactory{}
	}
	return singleInstance
}

// SetDefaultBackingStoreInstance allows configuring a custom backing store factory.
func SetDefaultBackingStoreInstance(store BackingStoreFactory) {
	lock.Lock()
	defer lock.Unlock()
	singleInstance = store
}
