package store

import (
	"sync"
)

var lock = &sync.Mutex{}

var singleInstance BackingStoreFactory

// GetInstance returns a backing store instance.
// if none exists an instance of inMemoryBackingStore is initialized and returned
func GetInstance() BackingStoreFactory {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		singleInstance = &InMemoryBackingStoreFactory{}
	}
	return singleInstance
}

// SetInstance allows configuring a custom backing store factory.
func SetInstance(store BackingStoreFactory) {
	lock.Lock()
	defer lock.Unlock()
	singleInstance = store
}
