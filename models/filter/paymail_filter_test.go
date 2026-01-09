package filter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPaymailFilter(t *testing.T) {
	t.Parallel()

	t.Run("default filter", func(t *testing.T) {
		filter := AdminPaymailFilter{}
		dbConditions := filter.ToDbConditions()

		assert.Len(t, dbConditions, 1)
		assert.Nil(t, dbConditions["deleted_at"])
	})

	t.Run("empty filter with include deleted", func(t *testing.T) {
		filter := fromJSON[AdminPaymailFilter](`{
			"includeDeleted": true
		}`)
		dbConditions := filter.ToDbConditions()

		assert.Empty(t, dbConditions)
	})

	t.Run("with alias", func(t *testing.T) {
		filter := fromJSON[AdminPaymailFilter](`{
			"alias": "example",
			"includeDeleted": true
		}`)
		dbConditions := filter.ToDbConditions()

		assert.Len(t, dbConditions, 1)
		assert.Equal(t, "example", dbConditions["alias"])
	})

	t.Run("with publicName", func(t *testing.T) {
		filter := fromJSON[AdminPaymailFilter](`{
			"publicName": "thepubname",
			"includeDeleted": true
		}`)
		dbConditions := filter.ToDbConditions()

		assert.Len(t, dbConditions, 1)
		assert.Equal(t, "thepubname", dbConditions["public_name"])
	})

	t.Run("with publicName", func(t *testing.T) {
		filter := fromJSON[AdminPaymailFilter](`{
			"publicName": "thepubname",
			"xpubId": "thexpubid",
			"includeDeleted": true
		}`)
		dbConditions := filter.ToDbConditions()

		assert.Len(t, dbConditions, 2)
		assert.Equal(t, "thexpubid", dbConditions["xpub_id"])
		assert.Equal(t, "thepubname", dbConditions["public_name"])
	})
}
