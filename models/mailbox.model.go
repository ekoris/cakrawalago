package models

type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by User to `profiles`
func (MailBox) TableName() string {
	return "mailbox"
}

type MailBox struct {
	ID     uint   `gorm:"primaryKey"`
	UserId string `gorm:"not null"`
	Name   string `gorm:"not null;unique"`
	Email  string `gorm:"not null"`
	Value  string `gorm:"not null"`
}
