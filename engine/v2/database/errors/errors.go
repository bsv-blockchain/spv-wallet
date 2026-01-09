//nolint:revive // Error types should be self-explanatory
package dberrors

import (
	"github.com/bsv-blockchain/spv-wallet/errdef"
	"github.com/joomcode/errorx"
)

var Namespace = errorx.NewNamespace("db")

var QueryFailed = Namespace.NewType("query_failed")

var NotFound = Namespace.NewType("not_found", errdef.TraitNotFound)
