package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"

	"github.com/vlad-zhogolev/credit-card-validator/validator"
)

type ValidationRequest struct {
	CreditCardNumber string `json:"credit_card_number"`
}

type ValidationResponse struct {
	IsValid bool `json:"is_valid"`
}

func main() {
	http.HandleFunc("/check_card_validity", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Handle card validity check request: %v\n", r.URL.RawQuery)
		query, err := url.QueryUnescape(r.URL.RawQuery)
		if err != nil {
			log.Printf("Failed to unescape query: %v\n", r.URL.RawQuery)
			w.WriteHeader(http.StatusNotFound)
			return
		}

		var validationRequest ValidationRequest
		err = json.Unmarshal([]byte(query), &validationRequest)
		if err != nil {
			log.Printf("Failed to parse request: %v\n", r.URL.RawQuery)
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}

		cardNumber, err := validator.NumbersFromString(validationRequest.CreditCardNumber)
		if err != nil {
			log.Printf("Failed convert card number: %v\n", validationRequest.CreditCardNumber)
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}

		data := ValidationResponse{IsValid: validator.Validate(cardNumber)}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
