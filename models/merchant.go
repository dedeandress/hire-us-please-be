package models

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Merchant struct {
	ID        *uuid.UUID      `gorm:"Type:uuid;NOT NULL;PRIMARY_KEY;DEFAULT:uuid_generate_v1()" json:"id" db:"id"`
	Name      string          `json:"name"`
	Longitude decimal.Decimal `gorm:"type:decimal(10,6);" json:"longitude"`
	Latitude  decimal.Decimal `gorm:"type:decimal(10,6);" json:"latitude"`
	Category  string          `json:"category"`
}
