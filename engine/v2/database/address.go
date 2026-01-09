package database

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"

	"github.com/bsv-blockchain/spv-wallet/models/bsv"
)

// Address represents a user's (bitcoin) addresses.
type Address struct {
	Address string `gorm:"type:char(34);primaryKey"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	CustomInstructions datatypes.JSONSlice[bsv.CustomInstruction]

	UserID string
	User   *User `gorm:"foreignKey:UserID"`
}
