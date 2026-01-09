package config

import (
	"github.com/bsv-blockchain/spv-wallet/engine/datastore"
	"github.com/bsv-blockchain/spv-wallet/engine/spverrors"
)

// Validate checks the configuration for specific rules
func (d *DbConfig) Validate() error {
	if d.Datastore.Engine == datastore.Empty || d.Datastore.Engine == "" {
		return spverrors.Newf("missing a valid datastore engine")
	}

	switch d.Datastore.Engine {
	case datastore.SQLite:
		if d.SQLite == nil {
			return spverrors.Newf("missing sqlite config")
		}
	case datastore.PostgreSQL:
		if d.SQL == nil {
			return spverrors.Newf("missing sql config")
		} else if len(d.SQL.Host) == 0 {
			return spverrors.Newf("missing sql host")
		} else if len(d.SQL.User) == 0 {
			return spverrors.Newf("missing sql username")
		} else if len(d.SQL.Name) == 0 {
			return spverrors.Newf("missing sql db name")
		}
	}

	return nil
}
