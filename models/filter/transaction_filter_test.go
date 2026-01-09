package filter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransactionFilter(t *testing.T) {
	t.Parallel()

	t.Run("default filter", func(t *testing.T) {
		filter := TransactionFilter{}
		dbConditions := filter.ToDbConditions()

		assert.Len(t, dbConditions, 1)
		assert.Nil(t, dbConditions["deleted_at"])
	})

	t.Run("empty filter with include deleted", func(t *testing.T) {
		filter := fromJSON[TransactionFilter](`{
			"includeDeleted": true
		}`)
		dbConditions := filter.ToDbConditions()

		assert.Empty(t, dbConditions)
	})

	t.Run("with hex", func(t *testing.T) {
		filter := fromJSON[TransactionFilter](`{
			"hex": "test",
			"includeDeleted": true
		}`)
		dbConditions := filter.ToDbConditions()

		assert.Len(t, dbConditions, 1)
		assert.Equal(t, "test", dbConditions["hex"])
	})

	t.Run("with block_height", func(t *testing.T) {
		filter := fromJSON[TransactionFilter](`{
			"blockHeight": 100,
			"includeDeleted": true
		}`)
		dbConditions := filter.ToDbConditions()

		assert.Len(t, dbConditions, 1)
		assert.Equal(t, uint64(100), dbConditions["block_height"])
	})
}
