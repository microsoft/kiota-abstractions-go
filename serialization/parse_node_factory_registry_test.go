package serialization

import (
	"testing"

	assert "github.com/stretchr/testify/assert"
)

func TestParseNodeFactoryRegistryHonoursInterface(t *testing.T) {
	assert.Implements(t, (*ParseNodeFactory)(nil), DefaultParseNodeFactoryInstance)
}
