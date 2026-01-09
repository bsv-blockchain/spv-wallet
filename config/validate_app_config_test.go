package config_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/bsv-blockchain/spv-wallet/config"
)

func TestValidateAppConfigForDefaultConfig(t *testing.T) {
	t.Parallel()

	// given:
	cfg := config.GetDefaultAppConfig()

	// when:
	err := cfg.Validate()

	// then:
	require.NoError(t, err)
}
