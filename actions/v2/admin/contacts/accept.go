package contacts

import (
	"net/http"

	"github.com/bsv-blockchain/spv-wallet/actions/v2/internal/mapping"
	"github.com/bsv-blockchain/spv-wallet/errdef/clienterr"
	"github.com/gin-gonic/gin"
)

// AdminAcceptInvitation accepts an invitation from a contact.
func (s *APIAdminContacts) AdminAcceptInvitation(c *gin.Context, id uint) {
	contact, err := s.contactsService.AcceptContactByID(c, id)
	if err != nil {
		clienterr.Response(c, err, s.logger)
		return
	}

	res := mapping.MapToContactContract(contact)

	c.JSON(http.StatusOK, res)
}
