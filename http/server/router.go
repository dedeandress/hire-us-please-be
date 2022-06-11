package server

import (
	"go_sample_login_register/http/server/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func buildRouteHandler() http.Handler {
	router := mux.NewRouter()
	routePost(router)

	return router
}

func routePost(router *mux.Router) {

	//auth
	router.HandleFunc("/v1/auth/login", handlers.HandleLogin).Methods(http.MethodPost)
	router.HandleFunc("/v1/auth/register", handlers.HandleRegister).Methods(http.MethodPost)
	router.HandleFunc("/v1/me", handlers.Auth(handlers.HandleGetMe)).Methods(http.MethodGet)

	//merchant
	router.HandleFunc("/v1/merchants", handlers.HandleGetMerchantList).Methods(http.MethodGet)
	router.HandleFunc("/v1/merchants/{merchantId}", handlers.HandleGetMerchantDetail).Methods(http.MethodGet)

}
