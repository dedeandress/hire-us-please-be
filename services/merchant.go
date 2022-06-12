package services

import (
	"go_sample_login_register/enums"
	"go_sample_login_register/models"
	"go_sample_login_register/params"
	"go_sample_login_register/repositories"
	"go_sample_login_register/validators"
	"strconv"
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


func GetFilteredMerchantList(latitude, longitude, maxDistance string) params.Response {
	var filteredMerchants []models.Merchant
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
	
	for _, merchant := range merchants {
		const bitSize = 64
		lat, err := strconv.ParseFloat(latitude, bitSize)
		if err != nil {
			return createResponseError(
				ResponseService{
					RollbackDB: true,
					Error:      err,
					ResultCode: enums.BAD_REQUEST,
				})
		}
		long, err := strconv.ParseFloat(longitude, bitSize)
		if err != nil {
			return createResponseError(
				ResponseService{
					RollbackDB: true,
					Error:      err,
					ResultCode: enums.BAD_REQUEST,
				})
		}
		md, err := strconv.ParseFloat(maxDistance, bitSize)
		if err != nil {
			return createResponseError(
				ResponseService{
					RollbackDB: true,
					Error:      err,
					ResultCode: enums.BAD_REQUEST,
				})
		}

		distance := validators.Location_distance_calc(lat, long, merchant.Latitude.InexactFloat64(), merchant.Longitude.InexactFloat64(), "K")
		if distance <= md {
			filteredMerchants = append(filteredMerchants, merchant)
		}
	}

	return createResponseSuccess(ResponseService{Payload: filteredMerchants})
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
