package contacts

import (
	"context"

	"github.com/rs/zerolog"

	"github.com/bsv-blockchain/spv-wallet/engine"
	"github.com/bsv-blockchain/spv-wallet/engine/v2/contacts/contactsmodels"
)

type contactsService interface {
	AcceptContact(ctx context.Context, userID, paymail string) error
	ConfirmContact(ctx context.Context, userID, paymail string) error
	Find(ctx context.Context, userID, paymail string) (*contactsmodels.Contact, error)
	RejectContact(ctx context.Context, userID, paymail string) error
	RemoveContact(ctx context.Context, userID, paymail string) error
	UnconfirmContact(ctx context.Context, userID, paymail string) error
	UpsertContact(ctx context.Context, newContact contactsmodels.NewContact) (*contactsmodels.Contact, error)
}

// APIContacts represents server with API endpoints
type APIContacts struct {
	contactsService contactsService
	logger          *zerolog.Logger
}

// NewAPIContacts creates a new server with API endpoints
func NewAPIContacts(engine engine.ClientInterface, log *zerolog.Logger) APIContacts {
	logger := log.With().Str("api", "contacts").Logger()

	return APIContacts{
		contactsService: engine.ContactService(),
		logger:          &logger,
	}
}
