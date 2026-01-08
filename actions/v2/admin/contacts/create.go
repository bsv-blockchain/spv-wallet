package contacts

import (
	"net/http"

	"github.com/bitcoin-sv/spv-wallet/actions/v2/internal/mapping"
	"github.com/bitcoin-sv/spv-wallet/api"
	"github.com/bitcoin-sv/spv-wallet/engine/v2/contacts/contactsmodels"
	"github.com/bitcoin-sv/spv-wallet/errdef/clienterr"
	"github.com/gin-gonic/gin"
)

// AdminCreateContact creates a new contact for a user.
func (s *APIAdminContacts) AdminCreateContact(c *gin.Context, paymail string) {
	var req api.RequestsAdminCreateContact
	if err := c.Bind(&req); err != nil {
		clienterr.UnprocessableEntity.New().Wrap(err).Response(c, s.logger)
		return
	}

	newContact := contactsmodels.NewContact{
		FullName:          req.FullName,
		NewContactPaymail: paymail,
		RequesterPaymail:  req.CreatorPaymail,
		Status:            contactsmodels.ContactNotConfirmed,
	}

	contact, err := s.contactsService.AdminCreateContact(c, newContact)
	if err != nil {
		clienterr.Response(c, err, s.logger)
		return
	}

	c.JSON(http.StatusOK, mapping.MapToContactContract(contact))
}
