package sql

// NOTE: Example tests for buildQueryForInputs were removed because GORM v1.30.0+ has a known
// issue (https://github.com/go-gorm/gorm/issues/5036) where ToSQL/DryRun mode doesn't correctly
// render nested subqueries passed to Table(). The actual SQL execution works correctly as
// verified by integration tests in selector_test.go.

import (
	"fmt"
	"time"

	"gorm.io/gorm"

	"github.com/bsv-blockchain/spv-wallet/engine/tester/tgorm"
	"github.com/bsv-blockchain/spv-wallet/models/bsv"
)

// ExampleUTXOSelector_buildUpdateTouchedAtQuery_sqlite demonstrates what would be the SQL statement used to update inputs after selecting them.
func ExampleUTXOSelector_buildUpdateTouchedAtQuery_sqlite() {
	db := tgorm.GormDBForPrintingSQL(tgorm.SQLite)

	selector := givenInputsSelector(db)

	utxos := []*selectedUTXO{
		{TxID: "tx_id_1", Vout: 0},
		{TxID: "tx_id_1", Vout: 1},
		{TxID: "tx_id_2", Vout: 0},
	}

	query := db.ToSQL(func(db *gorm.DB) *gorm.DB {
		query := selector.buildUpdateTouchedAtQuery(db, utxos)
		query.UpdateColumn("touched_at", time.Date(2006, 0o2, 0o1, 15, 4, 5, 7, time.UTC))
		return query
	})

	fmt.Println(query)

	// Output: UPDATE `xapi_user_utxos` SET `touched_at`="2006-02-01 15:04:05" WHERE (tx_id, vout) in (("tx_id_1",0),("tx_id_1",1),("tx_id_2",0))
}

// ExampleUTXOSelector_buildUpdateTouchedAtQuery_postgres demonstrates what would be the SQL statement used to update inputs after selecting them.
func ExampleUTXOSelector_buildUpdateTouchedAtQuery_postgres() {
	db := tgorm.GormDBForPrintingSQL(tgorm.PostgreSQL)

	selector := givenInputsSelector(db)

	utxos := []*selectedUTXO{
		{TxID: "tx_id_1", Vout: 0},
		{TxID: "tx_id_1", Vout: 1},
		{TxID: "tx_id_2", Vout: 0},
	}

	query := db.ToSQL(func(db *gorm.DB) *gorm.DB {
		query := selector.buildUpdateTouchedAtQuery(db, utxos)
		query.UpdateColumn("touched_at", time.Date(2006, 0o2, 0o1, 15, 4, 5, 7, time.UTC))
		return query
	})

	fmt.Println(query)

	// Output: UPDATE "xapi_user_utxos" SET "touched_at"='2006-02-01 15:04:05' WHERE (tx_id, vout) in (('tx_id_1',0),('tx_id_1',1),('tx_id_2',0))
}

func givenInputsSelector(db *gorm.DB) *UTXOSelector {
	selector := NewUTXOSelector(db, bsv.FeeUnit{Satoshis: 1, Bytes: 1000})
	return selector
}
