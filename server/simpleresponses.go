package server

import (
	"github.com/bsv-blockchain/spv-wallet/engine/spverrors"
	"github.com/gin-gonic/gin"
)

// NotFound handles all 404 requests
func NotFound(c *gin.Context) {
	spverrors.ErrorResponse(c, spverrors.ErrRouteNotFound, nil)
}

// MethodNotAllowed handles all 405 requests
func MethodNotAllowed(c *gin.Context) {
	spverrors.ErrorResponse(c, spverrors.ErrRouteMethodNotAllowed, nil)
}
