package taskmanager

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFactory_String(t *testing.T) {
	t.Run("test all factories", func(t *testing.T) {
		assert.Equal(t, "empty", FactoryEmpty.String())
		assert.Equal(t, "memory", FactoryMemory.String())
		assert.Equal(t, "redis", FactoryRedis.String())
	})
}

func TestFactory_IsEmpty(t *testing.T) {
	t.Run("test empty factory", func(t *testing.T) {
		f := FactoryEmpty
		assert.True(t, f.IsEmpty())
	})

	t.Run("test regular factory", func(t *testing.T) {
		f := FactoryMemory
		assert.False(t, f.IsEmpty())
	})
}
