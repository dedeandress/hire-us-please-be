package repositories

import (
	"go_sample_login_register/models"

	"github.com/jinzhu/gorm"
)

type MenuRepository interface {
	Insert(menu *models.Menu) (insertedMenu *models.Menu, err error)
	GetMenuByID(menuID string) (menu *models.Menu, err error)
	GetMenuByMerchantID(merchantId string) ([]models.Menu, error)
	Update(menu *models.Menu) (*models.Menu, error)
}

type menuRepository struct {
	db *DataSource
}

var menuRepo *menuRepository

func GetMenuRepository() MenuRepository {
	if DBTrx != nil {
		menuRepo = &menuRepository{db: DBTrx}
	} else {
		menuRepo = &menuRepository{db: DB}
	}
	return menuRepo
}

func (menuRepo *menuRepository) Insert(menu *models.Menu) (insertedMenu *models.Menu, err error) {
	insertedMenu = &models.Menu{}
	if err := menuRepo.db.Create(menu).Scan(insertedMenu).Error; err != nil {
		// fmt.Errorf("Error while inserting record in database (%s) ", err.Error())
		return nil, err
	}
	return insertedMenu, err
}

func (menuRepo *menuRepository) GetMenuByID(menuID string) (menu *models.Menu, err error) {
	menu = &models.Menu{}
	res := menuRepo.db.Scopes(filterMenuByID(menuID)).First(&menu)
	if err := res.Error; err != nil {
		return nil, err
	}
	return menu, err
}

func (menuRepo *menuRepository) GetMenuByMerchantID(merchantId string) ([]models.Menu, error) {
	menus := make([]models.Menu, 0)
	res := menuRepo.db.Scopes(filterMenuByMerchantID(merchantId)).Find(&menus)
	if err := res.Error; err != nil {
		return nil, err
	}
	return menus, nil
}

func (menuRepo *menuRepository) Update(menu *models.Menu) (*models.Menu, error) {
	if err := menuRepo.db.Model(&menu).Updates(menu).First(&menu).Error; err != nil {
		// fmt.Errorf("Error while updating record in database (%s) ", err.Error())
		return nil, err
	}
	return menu, nil
}

func filterMenuByID(menuID string) func(db *gorm.DB) *gorm.DB {
	return makeFilterFunc("menus.id = ?", menuID)
}

func filterMenuByMerchantID(merchantID string) func(db *gorm.DB) *gorm.DB {
	return makeFilterFunc("menus.merchant_id = ?", merchantID)
}
