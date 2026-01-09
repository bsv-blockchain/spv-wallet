package operations

import (
	"net/http"

	"github.com/bsv-blockchain/spv-wallet/actions/v2/operations/internal/mapping"
	"github.com/bsv-blockchain/spv-wallet/api"
	"github.com/bsv-blockchain/spv-wallet/engine/spverrors"
	"github.com/bsv-blockchain/spv-wallet/models/filter"
	"github.com/bsv-blockchain/spv-wallet/server/reqctx"
	"github.com/gin-gonic/gin"
)

// SearchOperations return operations based on given filter parameters
func (s *APIOperations) SearchOperations(c *gin.Context, params api.SearchOperationsParams) {
	userContext := reqctx.GetUserContext(c)
	userID, err := userContext.ShouldGetUserID()
	if err != nil {
		spverrors.AbortWithErrorResponse(c, err, s.logger)
		return
	}

	page := mapToFilter(params)
	pagedResult, err := s.engine.OperationsService().PaginatedForUser(c.Request.Context(), userID, page)
	if err != nil {
		spverrors.ErrorResponse(c, err, s.logger)
		return
	}

	c.JSON(http.StatusOK, mapping.OperationsPagedResponse(pagedResult))
}

func mapToFilter(params api.SearchOperationsParams) filter.Page {
	page := filter.Page{}

	if params.Page != nil {
		page.Number = *params.Page
	}
	if params.Size != nil {
		page.Size = *params.Size
	}
	if params.Sort != nil {
		page.Sort = *params.Sort
	}
	if params.SortBy != nil {
		page.SortBy = *params.SortBy
	}

	return page
}
