//nolint:revive // Traits should be self-explanatory
package errdef

import "github.com/joomcode/errorx"

var (
	TraitConfig            = errorx.RegisterTrait("config")
	TraitIllegalArgument   = errorx.RegisterTrait("illegal_argument")
	TraitNotFound          = errorx.RegisterTrait("not_found")
	TraitAuth              = errorx.RegisterTrait("auth")
	TraitARC               = errorx.RegisterTrait("arc")
	TraitShouldNeverHappen = errorx.RegisterTrait("should_never_happen")
	TraitUnsupported       = errorx.RegisterTrait("unsupported")
)
