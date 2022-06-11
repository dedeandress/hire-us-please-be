package params

import (
	"github.com/shopspring/decimal"
	"go_sample_login_register/models"
)

type MerchantResponse struct {
	ID        string          `json:"id"`
	Name      string          `json:"name"`
	Longitude decimal.Decimal `json:"longitude"`
	Latitude  decimal.Decimal `json:"latitude"`
	Category  string
	Menus     []models.Menu
}
