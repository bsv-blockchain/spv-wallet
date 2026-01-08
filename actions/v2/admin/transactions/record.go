package transactions

import (
	"github.com/bitcoin-sv/spv-wallet/actions/v2/admin/internal/mapping"
	"github.com/bitcoin-sv/spv-wallet/api"
	"github.com/bitcoin-sv/spv-wallet/errdef/clienterr"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// RecordTransactionOutlineForUser records transaction outline for given user
func (s *APIAdminTransactions) RecordTransactionOutlineForUser(c *gin.Context) {
	var requestBody api.RequestsRecordTransactionOutlineForUser
	err := c.ShouldBindWith(&requestBody, binding.JSON)
	if err != nil {
		clienterr.UnprocessableEntity.New().Wrap(err).Response(c, s.logger)
		return
	}
	if requestBody.UserID == "" {
		clienterr.BadRequest.Detailed("missing_user_id", "userID not provided").Response(c, s.logger)
		return
	}

	outline, err := mapping.RequestsTransactionOutlineToOutline(&requestBody)
	if err != nil {
		clienterr.Response(c, err, s.logger)
		return
	}

	recorded, err := s.transactionsRecordService.RecordTransactionOutline(c, requestBody.UserID, outline)
	if err != nil {
		clienterr.Response(c, err, s.logger)
		return
	}

	c.JSON(201, mapping.RecordedOutline(recorded))
}
