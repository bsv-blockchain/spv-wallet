package contacts

import (
	"net/http"

	"github.com/bitcoin-sv/spv-wallet/errdef/clienterr"
	"github.com/gin-gonic/gin"
)

// AdminRejectInvitation rejects an invitation from a contact.
func (s *APIAdminContacts) AdminRejectInvitation(c *gin.Context, id uint) {
	_, err := s.contactsService.RejectContactByID(c, id)
	if err != nil {
		clienterr.Response(c, err, s.logger)
		return
	}

	c.Status(http.StatusOK)
}
