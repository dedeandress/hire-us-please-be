package handlers

import (
	"go_sample_login_register/services"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleGetMerchantList(w http.ResponseWriter, r *http.Request) {
	latitude := r.URL.Query().Get("latitude")
	longitude := r.URL.Query().Get("longitude")
	maxDistance := r.URL.Query().Get("max_distance")
	if maxDistance == "" {
		maxDistance = "10"
	}

	if latitude == "" || longitude == "" {
		ToJSON(w, http.StatusBadRequest, badRequestResponse)
		return
	}

	response := services.GetMerchantList(latitude, longitude, maxDistance)
	ToJSON(w, response.HttpCode, response)
}

func HandleGetMerchantDetail(w http.ResponseWriter, r *http.Request) {
	muxParams := mux.Vars(r)
	merchantID := muxParams["merchantId"]

	response := services.GetMerchantDetail(merchantID)
	ToJSON(w, response.HttpCode, response)
}
