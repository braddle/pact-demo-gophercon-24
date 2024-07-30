package producer

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type HealthStatus struct {
	Status string `json:"status"`
}

type IceCream struct {
	ID                  string       `json:"id"`
	Barcode             string       `json:"barcode"`
	Name                string       `json:"name"`
	Manufacturer        Manufacturer `json:"manufacturer"`
	Ingredients         []string     `json:"ingredients"`
	Calories            int64        `json:"calories"`
	RecyclablePackaging bool         `json:"recyclable_packaging"`
	Rating              float64      `json:"rating"`
	Images              []Image      `json:"images"`
}

type Image struct {
	URL    string `json:"url"`
	Width  int64  `json:"width"`
	Height int64  `json:"height"`
}

type Manufacturer struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

func GetRouter() *mux.Router {
	r := mux.NewRouter()
	r.Use(loggingMiddleware)

	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		s := HealthStatus{Status: "OKKKK"}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(s)
	})

	r.HandleFunc("/icecream/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		ic := IceCream{
			ID:      vars["id"],
			Barcode: "12345678",
			Name:    "Orange Calippo",
			Manufacturer: Manufacturer{
				ID:      "walls",
				Name:    "Walls",
				Address: "An Office, Long Road, London, N1 3RE",
			},
			Ingredients:         []string{"Water", "orange juice", "sugar", "glucose syrup", "apple juice", "fructose syrup"},
			Calories:            434,
			RecyclablePackaging: false,
			Rating:              4.3,
			Images: []Image{
				{
					URL:    "https://www.donaldscreamices.co.uk/image/cache//catalog/Walls%20Ice%20Cream/Impulse/download%20(1)-800x600h.jpg",
					Width:  800,
					Height: 600,
				},
			},
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(ic)
	})

	return r
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
