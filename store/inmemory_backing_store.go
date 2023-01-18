package store

import (
	"errors"
	"github.com/google/uuid"
	"reflect"
	"strings"
)

type BackingStoreSubscriber func(key string, oldVal interface{}, newVal interface{})

type inMemoryBackingStore struct {
	returnOnlyChangedValues bool
	initializationCompleted bool
	store                   map[string]interface{}
	subscribers             map[string]BackingStoreSubscriber
	changedValues           map[string]bool
}

// NewInMemoryBackingStore returns a new instance of an in memory backing store
func NewInMemoryBackingStore() BackingStore {
	return &inMemoryBackingStore{
		returnOnlyChangedValues: false,
		initializationCompleted: true,
		store:                   make(map[string]interface{}),
		subscribers:             make(map[string]BackingStoreSubscriber),
		changedValues:           make(map[string]bool),
	}
}

func (i *inMemoryBackingStore) Get(key string) (interface{}, error) {
	key = strings.TrimSpace(key)
	if key == "" {
		return nil, errors.New("key cannot be an empty string")
	}

	objectVal := i.store[key]

	if (i.GetReturnOnlyChangedValues() && i.changedValues[key]) || !i.GetReturnOnlyChangedValues() {
		return objectVal, nil
	} else {
		return nil, nil
	}
}

func (i *inMemoryBackingStore) Set(key string, value interface{}) error {
	key = strings.TrimSpace(key)
	if key == "" {
		return errors.New("key cannot be an empty string")
	}

	current := i.store[key]

	// check if objects values have changed
	if current == nil || !reflect.DeepEqual(current, value) {
		// track changed key
		i.changedValues[key] = i.GetInitializationCompleted()

		// update changed values
		i.store[key] = value

		// notify subs
		for _, subscriber := range i.subscribers {
			subscriber(key, current, value)
		}
	}

	return nil
}

func (i *inMemoryBackingStore) Enumerate() map[string]interface{} {
	items := make(map[string]interface{})

	for k, v := range i.store {
		if !i.GetReturnOnlyChangedValues() || i.changedValues[k] { // change flag not set or object changed
			items[k] = v
		}
	}

	return items
}

func (i *inMemoryBackingStore) EnumerateKeysForValuesChangedToNil() []string {
	keys := make([]string, 0)
	for k, v := range i.store {
		if i.changedValues[k] && v == nil {
			keys = append(keys, k)
		}
	}

	return keys
}

func (i *inMemoryBackingStore) Subscribe(callback BackingStoreSubscriber) string {
	id := uuid.New().String()
	i.subscribers[id] = callback
	return id
}

func (i *inMemoryBackingStore) SubscribeWithId(callback BackingStoreSubscriber, subscriptionId string) error {
	subscriptionId = strings.TrimSpace(subscriptionId)
	if subscriptionId == "" {
		return errors.New("subscriptionId cannot be an empty string")
	}

	i.subscribers[subscriptionId] = callback

	return nil
}

func (i *inMemoryBackingStore) Unsubscribe(subscriptionId string) error {
	subscriptionId = strings.TrimSpace(subscriptionId)
	if subscriptionId == "" {
		return errors.New("subscriptionId cannot be an empty string")
	}

	delete(i.subscribers, subscriptionId)

	return nil
}

func (i *inMemoryBackingStore) Clear() {
	for k := range i.store {
		delete(i.store, k)
		delete(i.changedValues, k) // changed values must be an element in the store
	}
}

func (i *inMemoryBackingStore) GetInitializationCompleted() bool {
	return i.initializationCompleted
}

func (i *inMemoryBackingStore) SetInitializationCompleted(val bool) {
	i.initializationCompleted = val
}

func (i *inMemoryBackingStore) GetReturnOnlyChangedValues() bool {
	return i.returnOnlyChangedValues
}

func (i *inMemoryBackingStore) SetReturnOnlyChangedValues(val bool) {
	i.returnOnlyChangedValues = val
}
