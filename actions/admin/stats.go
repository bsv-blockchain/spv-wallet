package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/bsv-blockchain/spv-wallet/engine/spverrors"
	"github.com/bsv-blockchain/spv-wallet/mappings"
	"github.com/bsv-blockchain/spv-wallet/server/reqctx"
)

// @Summary			Get stats
// @Description		Get statistics of the spv-wallet
// @Tags			Admin
// @Produce			json
// @Success			200	{object} response.AdminStats "Stats for the admin"
// @Failure 		500	"Internal Server Error - Error while fetching admin stats"
// @Router			/api/v1/admin/stats [get]
// @Security		x-auth-xpub
func stats(c *gin.Context, _ *reqctx.AdminContext) {
	stats, err := reqctx.Engine(c).GetStats(c.Request.Context())
	if err != nil {
		spverrors.ErrorResponse(c, err, reqctx.Logger(c))
		return
	}

	contract := mappings.MapToAdminStatsContract(stats)
	c.JSON(http.StatusOK, contract)
}
