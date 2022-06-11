package models

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Menu struct {
	ID         *uuid.UUID      `gorm:"Type:uuid;NOT NULL;PRIMARY_KEY;DEFAULT:uuid_generate_v1()" json:"id" db:"id"`
	Name       string          `json:"name"`
	Image      string          `json:"image"`
	Price      decimal.Decimal `gorm:"type:decimal(20,2);" json:"price"`
	MerchantID uuid.UUID       `json:"merchant_id"`
}
