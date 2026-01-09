package data

import (
	"net/http"

	"github.com/bsv-blockchain/spv-wallet/actions/v2/data/internal/mapping"
	"github.com/bsv-blockchain/spv-wallet/engine/spverrors"
	"github.com/bsv-blockchain/spv-wallet/models/bsv"
	"github.com/bsv-blockchain/spv-wallet/server/reqctx"
	"github.com/gin-gonic/gin"
)

// DataById returns data for user by its id
func (s *APIData) DataById(c *gin.Context, id string) {
	userContext := reqctx.GetUserContext(c)
	userID, err := userContext.ShouldGetUserID()
	if err != nil {
		spverrors.ErrorResponse(c, err, s.logger)
		return
	}

	_, err = bsv.OutpointFromString(id)
	if err != nil {
		spverrors.ErrorResponse(c, spverrors.ErrInvalidDataID.Wrap(err), s.logger)
		return
	}

	data, err := s.engine.DataService().FindForUser(c.Request.Context(), id, userID)
	if err != nil {
		spverrors.ErrorResponse(c, err, s.logger)
		return
	}

	if data == nil {
		spverrors.ErrorResponse(c, spverrors.ErrDataNotFound, s.logger)
		return
	}

	c.JSON(http.StatusOK, mapping.DataResponse(data))
}
