package services

import (
	"go_sample_login_register/enums"
	"go_sample_login_register/params"
	"go_sample_login_register/repositories"
)

func GetMerchantList() params.Response {
	merchantRepo := repositories.GetMerchantRepository()

	merchants, err := merchantRepo.GetMerchantList()
	if err != nil {
		return createResponseError(
			ResponseService{
				RollbackDB: true,
				Error:      err,
				ResultCode: enums.BAD_REQUEST,
			})
	}

	return createResponseSuccess(ResponseService{Payload: merchants})
}

func GetMerchantDetail(merchantID string) params.Response {
	merchantRepo := repositories.GetMerchantRepository()
	menuRepo := repositories.GetMenuRepository()

	merchant, err := merchantRepo.GetMerchantByID(merchantID)
	if err != nil {
		return createResponseError(
			ResponseService{
				RollbackDB: true,
				Error:      err,
				ResultCode: enums.BAD_REQUEST,
			})
	}

	menu, err := menuRepo.GetMenuByMerchantID(merchantID)
	if err != nil {
		return createResponseError(
			ResponseService{
				RollbackDB: true,
				Error:      err,
				ResultCode: enums.BAD_REQUEST,
			})
	}

	merchantResponse := params.MerchantResponse{
		ID:        merchant.ID.String(),
		Name:      merchant.Name,
		Latitude:  merchant.Latitude,
		Longitude: merchant.Longitude,
		Category:  merchant.Category,
		Menus:     menu,
	}

	return createResponseSuccess(ResponseService{Payload: merchantResponse})
}
