package filter

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimeRange(t *testing.T) {
	t.Parallel()

	t.Run("empty time range", func(t *testing.T) {
		filter := TimeRange{}
		dbConditions := filter.ToDbConditions()

		assert.True(t, filter.isEmpty())
		assert.Empty(t, dbConditions)
	})

	t.Run("only _from_ field", func(t *testing.T) {
		timeNow := time.Now()
		filter := TimeRange{
			From: ptr(timeNow),
		}
		dbConditions := filter.ToDbConditions()

		assert.Len(t, dbConditions, 1)
		assert.Equal(t, timeNow, dbConditions["$gte"])
	})

	t.Run("only _to_ field", func(t *testing.T) {
		timeNow := time.Now()
		filter := TimeRange{
			To: ptr(timeNow),
		}
		dbConditions := filter.ToDbConditions()

		assert.Len(t, dbConditions, 1)
		assert.Equal(t, timeNow, dbConditions["$lte"])
	})
}
