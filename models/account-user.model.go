package models

import (
	"time"
)

type TableAccount interface {
	TableName() string
}

// TableName overrides the table name used by User to `profiles`
func (Account) TableAccount() string {
	return "accounts"
}

type Account struct {
	ID                 int64 `gorm:"primaryKey" json:"id"`
	UserId             int64
	Name               string
	DateOfBirth        time.Time
	PlaceOfBirth       string
	Nik                int64
	Address            string
	MarketId           int64
	IdentityAttachment string
	SelfPhoto          string
	SignaturePhoto     string
	CreatedAt          time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt          time.Time `gorm:"not null" json:"updated_at"`
}
