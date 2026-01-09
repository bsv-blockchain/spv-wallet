package contacts

import (
	"net/http"

	"github.com/bsv-blockchain/spv-wallet/api"
	"github.com/bsv-blockchain/spv-wallet/errdef/clienterr"
	"github.com/gin-gonic/gin"
)

// AdminConfirmContact confirms a contact between two users.
func (s *APIAdminContacts) AdminConfirmContact(c *gin.Context) {
	var reqParams *api.RequestsAdminConfirmContact
	if err := c.Bind(&reqParams); err != nil {
		clienterr.UnprocessableEntity.New().Wrap(err).Response(c, s.logger)
		return
	}

	if err := s.contactsService.AdminConfirmContacts(c, reqParams.PaymailA, reqParams.PaymailB); err != nil {
		clienterr.Response(c, err, s.logger)
		return
	}

	c.Status(http.StatusOK)

}
