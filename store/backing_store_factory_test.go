package store

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetDefaultBackingStoreInstance(t *testing.T) {
	assert.NotNil(t, BackingStoreFactoryInstance)
}
