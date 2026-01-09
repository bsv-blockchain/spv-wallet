package contacts

import (
	"net/http"

	"github.com/bsv-blockchain/spv-wallet/errdef/clienterr"
	"github.com/gin-gonic/gin"
)

// AdminDeleteContact deletes a contact
func (s *APIAdminContacts) AdminDeleteContact(c *gin.Context, id uint) {
	err := s.contactsService.RemoveContactByID(c, id)
	if err != nil {
		clienterr.Response(c, err, s.logger)
		return
	}

	c.Status(http.StatusOK)
}
