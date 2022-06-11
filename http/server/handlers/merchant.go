package handlers

import (
	"github.com/gorilla/mux"
	"go_sample_login_register/services"
	"net/http"
)

func HandleGetMerchantList(w http.ResponseWriter, r *http.Request) {

	response := services.GetMerchantList()
	ToJSON(w, response.HttpCode, response)
}

func HandleGetMerchantDetail(w http.ResponseWriter, r *http.Request) {
	muxParams := mux.Vars(r)
	merchantID := muxParams["merchantId"]

	response := services.GetMerchantDetail(merchantID)
	ToJSON(w, response.HttpCode, response)
}
