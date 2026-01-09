package paymail

import (
	"context"

	"github.com/bsv-blockchain/go-paymail/spv"
	primitives "github.com/bsv-blockchain/go-sdk/primitives/ec"
	trx "github.com/bsv-blockchain/go-sdk/transaction"
	"github.com/bsv-blockchain/spv-wallet/engine/v2/addresses/addressesmodels"
	"github.com/bsv-blockchain/spv-wallet/engine/v2/contacts/contactsmodels"
	"github.com/bsv-blockchain/spv-wallet/engine/v2/paymails/paymailsmodels"
)

// PaymailsService is an interface for paymails service
type PaymailsService interface {
	Find(ctx context.Context, alias, domain string) (*paymailsmodels.Paymail, error)
}

// UsersService is an interface for users service
type UsersService interface {
	GetPubKey(ctx context.Context, userID string) (*primitives.PublicKey, error)
}

// AddressesService is an interface for addresses service
type AddressesService interface {
	Create(ctx context.Context, newAddress *addressesmodels.NewAddress) error
}

// MerkleRootsVerifier is an interface for verifying merkle roots
type MerkleRootsVerifier interface {
	VerifyMerkleRoots(ctx context.Context, merkleRoots []*spv.MerkleRootConfirmationRequestItem) (bool, error)
}

// TxRecorder is an interface for recording transactions
type TxRecorder interface {
	RecordPaymailTransaction(ctx context.Context, tx *trx.Transaction, senderPaymail, receiverPaymail string) error
}

// ContactsService is an interface for contacts service
type ContactsService interface {
	AddContactRequest(ctx context.Context, fullName, paymail, userID string) (*contactsmodels.Contact, error)
}
