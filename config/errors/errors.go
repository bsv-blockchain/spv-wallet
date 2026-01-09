//nolint:revive // Error types should be self-explanatory
package errors

import (
	"github.com/joomcode/errorx"

	"github.com/bsv-blockchain/spv-wallet/errdef"
)

var Namespace = errorx.NewNamespace("config", errdef.TraitConfig)

var UnsupportedDomain = Namespace.NewType("unsupported_domain", errdef.TraitIllegalArgument)
