package filter

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUtxoFilter(t *testing.T) {
	t.Parallel()

	t.Run("default filter", func(t *testing.T) {
		filter := UtxoFilter{}
		dbConditions, err := filter.ToDbConditions()

		require.NoError(t, err)
		assert.Len(t, dbConditions, 1)
		assert.Nil(t, dbConditions["deleted_at"])
	})

	t.Run("empty filter with include deleted", func(t *testing.T) {
		filter := fromJSON[UtxoFilter](`{
			"includeDeleted": true
		}`)
		dbConditions, err := filter.ToDbConditions()

		require.NoError(t, err)
		assert.Empty(t, dbConditions)
	})

	t.Run("with type", func(t *testing.T) {
		filter := fromJSON[UtxoFilter](`{
			"type": "pubkey",
			"includeDeleted": true
		}`)
		dbConditions, err := filter.ToDbConditions()

		require.NoError(t, err)
		assert.Len(t, dbConditions, 1)
		assert.Equal(t, "pubkey", dbConditions["type"])
	})

	t.Run("with wrong type", func(t *testing.T) {
		filter := fromJSON[UtxoFilter](`{
			"type": "wrong_type",
			"includeDeleted": true
		}`)
		dbConditions, err := filter.ToDbConditions()

		require.Error(t, err)
		assert.Nil(t, dbConditions)
	})

	t.Run("admin filter with xpubid", func(t *testing.T) {
		filter := fromJSON[AdminUtxoFilter](`{
			"includeDeleted": true,
			"id": "theid",
			"xpubId": "thexpubid"
		}`)
		dbConditions, _ := filter.ToDbConditions()

		assert.Equal(t, "thexpubid", dbConditions["xpub_id"])
		assert.Equal(t, "theid", dbConditions["id"])
	})
}
