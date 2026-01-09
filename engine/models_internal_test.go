package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModelSetRecordTime(t *testing.T) {
	t.Parallel()

	t.Run("empty model", func(t *testing.T) {
		m := new(Model)
		assert.True(t, m.CreatedAt.IsZero())
		assert.True(t, m.UpdatedAt.IsZero())
	})

	t.Run("set created at time", func(t *testing.T) {
		m := new(Model)
		m.SetRecordTime(true)
		assert.False(t, m.CreatedAt.IsZero())
		assert.True(t, m.UpdatedAt.IsZero())
	})

	t.Run("set updated at time", func(t *testing.T) {
		m := new(Model)
		m.SetRecordTime(false)
		assert.True(t, m.CreatedAt.IsZero())
		assert.False(t, m.UpdatedAt.IsZero())
	})

	t.Run("set both times", func(t *testing.T) {
		m := new(Model)
		m.SetRecordTime(false)
		m.SetRecordTime(true)
		assert.False(t, m.CreatedAt.IsZero())
		assert.False(t, m.UpdatedAt.IsZero())
	})
}

func TestModelNew(t *testing.T) {
	t.Parallel()

	t.Run("New model", func(t *testing.T) {
		m := new(Model)
		assert.False(t, m.IsNew())
	})

	t.Run("set New flag", func(t *testing.T) {
		m := new(Model)
		m.New()
		assert.True(t, m.IsNew())
	})
}

func TestModelGetOptions(t *testing.T) {
	// t.Parallel()

	t.Run("base model", func(t *testing.T) {
		m := new(Model)
		opts := m.GetOptions(false)
		assert.Empty(t, opts)
	})

	t.Run("new record model", func(t *testing.T) {
		m := new(Model)
		opts := m.GetOptions(true)
		assert.Len(t, opts, 1)
	})
}

func TestModel_IsNew(t *testing.T) {
	t.Parallel()

	t.Run("base model", func(t *testing.T) {
		m := new(Model)
		assert.False(t, m.IsNew())
	})

	t.Run("New model", func(t *testing.T) {
		m := new(Model)
		m.New()
		assert.True(t, m.IsNew())
	})
}

func TestModel_RawXpub(t *testing.T) {
	m := new(Model)
	m.rawXpubKey = "xpub661MyMwAqRbcFqp1qzrF2AryEo4X8W1CNSAiT7t2wgXxkbt8nSrdZFYQeD19aTeiPtpAHDGtNUBxgFAg5d2GMzbAiVEsP8DJPLgTQ2LvZTz"
	assert.Equal(t, "xpub661MyMwAqRbcFqp1qzrF2AryEo4X8W1CNSAiT7t2wgXxkbt8nSrdZFYQeD19aTeiPtpAHDGtNUBxgFAg5d2GMzbAiVEsP8DJPLgTQ2LvZTz", m.RawXpub())
}

func TestModel_Name(t *testing.T) {
	t.Parallel()

	t.Run("base model", func(t *testing.T) {
		m := new(Model)
		assert.Empty(t, m.Name())
	})

	t.Run("set model name", func(t *testing.T) {
		m := new(Model)
		m.name = ModelXPub
		assert.Equal(t, "xpub", m.Name())
	})
}
