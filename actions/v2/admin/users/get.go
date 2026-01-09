package users

import (
	"net/http"

	"github.com/bsv-blockchain/spv-wallet/actions/v2/admin/internal/mapping"
	dberrors "github.com/bsv-blockchain/spv-wallet/engine/v2/database/errors"
	"github.com/bsv-blockchain/spv-wallet/errdef/clienterr"
	"github.com/gin-gonic/gin"
)

// UserById returns a user by ID
func (s *APIAdminUsers) UserById(c *gin.Context, id string) {
	user, err := s.engine.UsersService().GetByID(c, id)
	if err != nil {
		clienterr.Map(err).
			IfOfType(dberrors.NotFound).Then(clienterr.NotFound.New()).
			Response(c, s.logger)
		return
	}

	c.JSON(http.StatusOK, mapping.UserToResponse(user))
}
