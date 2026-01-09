//nolint:revive // Error types should be self-explanatory
package paymailerrors

import (
	"github.com/joomcode/errorx"

	"github.com/bsv-blockchain/spv-wallet/errdef"
)

var Namespace = errorx.NewNamespace("paymail")

var (
	InvalidAvatarURL        = Namespace.NewType("invalid_avatar_url", errdef.TraitIllegalArgument)
	InvalidPaymailAddress   = Namespace.NewType("invalid_paymail_address", errdef.TraitIllegalArgument)
	UserDoesntExist         = Namespace.NewType("user_doesnt_exist", errdef.TraitNotFound)
	NoDefaultPaymailAddress = Namespace.NewType("no_default_paymail_address", errdef.TraitIllegalArgument)
)
