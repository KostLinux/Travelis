package middleware

import (
	"io"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func GetSearchLocations(w http.ResponseWriter, r *http.Request) {
	locale := r.URL.Query().Get("locale")
	name := r.URL.Query().Get("name")

	// Validate query params
	if locale == "" {
		StatusBadRequest(w, "locale is required")
		return
	}

	if name == "" {
		StatusBadRequest(w, "name is required")
		return
	}

	bookingRapidAPIUrl := "https://booking-com.p.rapidapi.com/v1/hotels/locations?locale=" + locale + "&name=" + name

	req, err := http.NewRequest("GET", bookingRapidAPIUrl, nil)
	if err != nil {
		log.Printf("unable to create request: %v", err)
		StatusInternalServerError(w)
		return
	}

	req.Header.Add("X-RapidAPI-Key", os.Getenv("RAPIDAPI_KEY"))
	req.Header.Add("X-RapidAPI-Host", "booking-com.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("unable to get hotel locations: %v", err)
		StatusInternalServerError(w)
		return
	}

	if res.StatusCode != http.StatusOK {
		log.Printf("unexpected API status code: %v", res.StatusCode)
		StatusInternalServerError(w)
		return
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("unable to read response body: %v", err)
		StatusInternalServerError(w)
		return
	}

	StatusOK(w, body)
}

// SkyScanner
func GetCultureMarkets(w http.ResponseWriter, r *http.Request) {
	locale := r.URL.Query().Get("locale")

	// Validate query params
	if locale == "" {
		StatusBadRequest(w, "locale is required")
		return
	}

	skyscannerAPIUrl := "https://partners.api.skyscanner.net/apiservices/v3/culture/markets/" + locale

	req, err := http.NewRequest("GET", skyscannerAPIUrl, nil)
	if err != nil {
		log.Printf("unable to create request: %v", err)
		StatusInternalServerError(w)
		return
	}

	req.Header.Add("x-api-key", os.Getenv("SKYSCANNER_API_KEY"))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("unable to get markets locations: %v", err)
		StatusInternalServerError(w)
		return
	}

	if res.StatusCode != http.StatusOK {
		log.Printf("unexpected API status code: %v", res.StatusCode)
		StatusInternalServerError(w)
		return
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("unable to read response body: %v", err)
		StatusInternalServerError(w)
		return
	}

	StatusOK(w, body)
}
