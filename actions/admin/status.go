package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/bsv-blockchain/spv-wallet/server/reqctx"
)

// @Summary			Get status
// @Description		Get status
// @Tags			Admin
// @Produce			json
// @Success			200 {boolean} bool "Status response"
// @Router			/api/v1/admin/status [get]
// @Security		x-auth-xpub
func status(c *gin.Context, _ *reqctx.AdminContext) {
	c.JSON(http.StatusOK, true)
}
