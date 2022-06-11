package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/kr/pretty"
	"go_sample_login_register/models"
	"log"
)

type MerchantRepository interface {
	Insert(merchant *models.Merchant) (insertedMerchant *models.Merchant, err error)
	GetMerchantByID(merchantID string) (merchant *models.Merchant, err error)
	Update(merchant *models.Merchant) (*models.Merchant, error)
	GetMerchantList() ([]models.Merchant, error)
}

type merchantRepository struct {
	db *DataSource
}

var merchantRepo *merchantRepository

func GetMerchantRepository() MerchantRepository {
	if DBTrx != nil {
		merchantRepo = &merchantRepository{db: DBTrx}
	} else {
		merchantRepo = &merchantRepository{db: DB}
	}
	return merchantRepo
}

func (merchantRepo *merchantRepository) Insert(merchant *models.Merchant) (insertedMerchant *models.Merchant, err error) {
	insertedMerchant = &models.Merchant{}
	if err := merchantRepo.db.Create(merchant).Scan(insertedMerchant).Error; err != nil {
		// fmt.Errorf("Error while inserting record in database (%s) ", err.Error())
		return nil, err
	}
	return insertedMerchant, err
}

func (merchantRepo *merchantRepository) GetMerchantList() ([]models.Merchant, error) {
	var merchantList []models.Merchant
	res := merchantRepo.db.Find(&merchantList)
	if res.Error != nil {
		return nil, res.Error
	}

	log.Print(pretty.Sprint(merchantList))
	return merchantList, nil
}

func (merchantRepo *merchantRepository) GetMerchantByID(merchantID string) (merchant *models.Merchant, err error) {
	merchant = &models.Merchant{}
	res := merchantRepo.db.Scopes(filterMerchantByID(merchantID)).First(&merchant)
	if err := res.Error; err != nil {
		// fmt.Errorf("Error querying for user for user with ID %s and error: %s", userID, err.Error())
		return nil, err
	}
	return merchant, err
}

func (merchantRepo *merchantRepository) Update(merchant *models.Merchant) (*models.Merchant, error) {
	if err := merchantRepo.db.Model(&merchant).Updates(merchant).First(&merchant).Error; err != nil {
		// fmt.Errorf("Error while updating record in database (%s) ", err.Error())
		return nil, err
	}
	return merchant, nil
}

func filterMerchantByID(merchantID string) func(db *gorm.DB) *gorm.DB {
	return makeFilterFunc("merchants.id = ?", merchantID)
}
