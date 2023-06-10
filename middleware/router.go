package middleware

import (
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	// Search locations from Booking.com
	router.HandleFunc("/v1/search-locations", GetSearchLocations).Methods("GET")

	return router
}
