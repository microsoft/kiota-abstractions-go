package store

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetDefaultBackingStoreInstance(t *testing.T) {
	assert.NotNil(t, GetDefaultBackingStoreInstance())
}

func TestSetDefaultBackingStoreInstance(t *testing.T) {

	testFactory := testFactory{}

	assert.NotEqual(t, GetDefaultBackingStoreInstance(), &testFactory)

	SetDefaultBackingStoreInstance(&testFactory)

	assert.Equal(t, GetDefaultBackingStoreInstance(), &testFactory)
}

type testFactory struct {
}

func (i *testFactory) CreateBackingStore() BackingStore {
	return NewInMemoryBackingStore()
}
