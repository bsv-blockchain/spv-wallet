package filter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccessKeyFilter(t *testing.T) {
	t.Parallel()

	t.Run("default filter", func(t *testing.T) {
		filter := AccessKeyFilter{}
		dbConditions := filter.ToDbConditions()

		assert.Len(t, dbConditions, 1)
		assert.Nil(t, dbConditions["deleted_at"])
	})

	t.Run("empty filter with include deleted", func(t *testing.T) {
		filter := fromJSON[AccessKeyFilter](`{
			"includeDeleted": true
		}`)
		dbConditions := filter.ToDbConditions()

		assert.Empty(t, dbConditions)
	})

	t.Run("with full RevokedRange", func(t *testing.T) {
		filter := fromJSON[AccessKeyFilter](`{
			"includeDeleted": true,
			"revokedRange": {
				"from": "2024-02-26T11:01:28Z",
				"to": "2024-02-25T11:01:28Z"
			}
		}`)
		dbConditions := filter.ToDbConditions()

		assert.Len(t, dbConditions["revoked_at"].(map[string]interface{}), 2)
	})

	t.Run("with empty RevokedRange", func(t *testing.T) {
		filter := fromJSON[AccessKeyFilter](`{
			"includeDeleted": true,
			"revokedRange": {}
		}`)
		dbConditions := filter.ToDbConditions()

		assert.Empty(t, dbConditions)
	})

	t.Run("with partially filled RevokedRange", func(t *testing.T) {
		filter := fromJSON[AccessKeyFilter](`{
			"includeDeleted": true,
			"revokedRange": {
				"from": "2024-02-26T11:01:28Z"
			}
		}`)
		dbConditions := filter.ToDbConditions()

		assert.Len(t, dbConditions["revoked_at"].(map[string]interface{}), 1)
	})

	t.Run("admin filter with xpubid", func(t *testing.T) {
		filter := fromJSON[AdminAccessKeyFilter](`{
			"includeDeleted": true,
			"xpubId": "thexpubid"
		}`)
		dbConditions := filter.ToDbConditions()

		assert.Equal(t, "thexpubid", dbConditions["xpub_id"])
	})
}
