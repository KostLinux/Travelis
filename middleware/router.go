package middleware

import (
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	// Search locations from Booking.com
	router.HandleFunc("/v1/search-locations", GetSearchLocations).Methods("GET")

	// SkyScanner Culture Marketes
	router.HandleFunc("/v1/get-culture-markets", GetCultureMarkets).Methods("GET")
	return router
}
