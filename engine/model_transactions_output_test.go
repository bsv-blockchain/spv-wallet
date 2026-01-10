package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestXpubOutputValue_Scan(t *testing.T) {
	t.Parallel()

	t.Run("nil value", func(t *testing.T) {
		x := XpubOutputValue{}
		err := x.Scan(nil)
		require.NoError(t, err)
		assert.Empty(t, x)
	})

	t.Run("empty string", func(t *testing.T) {
		x := XpubOutputValue{}
		err := x.Scan([]byte("\"\""))
		require.NoError(t, err)
		assert.Empty(t, x)
	})

	t.Run("empty string - incorrectly coded", func(t *testing.T) {
		x := XpubOutputValue{}
		err := x.Scan([]byte(""))
		require.NoError(t, err)
		assert.Empty(t, x)
	})

	t.Run("object", func(t *testing.T) {
		x := XpubOutputValue{}
		err := x.Scan([]byte("{\"xPubId\":543}"))
		require.NoError(t, err)
		assert.Len(t, x, 1)
		assert.Equal(t, int64(543), x["xPubId"])
	})
}

func TestXpubOutputValue_Value(t *testing.T) {
	t.Parallel()

	t.Run("empty object", func(t *testing.T) {
		x := XpubOutputValue{}
		value, err := x.Value()
		require.NoError(t, err)
		assert.Equal(t, "{}", value)
	})

	t.Run("map present", func(t *testing.T) {
		x := XpubOutputValue{
			"xPubId": 123,
		}
		value, err := x.Value()
		require.NoError(t, err)
		assert.Equal(t, "{\"xPubId\":123}", value)
	})
}
