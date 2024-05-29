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
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		query, err := url.QueryUnescape(r.URL.RawQuery)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		var validationRequest ValidationRequest
		err = json.Unmarshal([]byte(query), &validationRequest)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		cardNumber, _ := validator.NumbersFromString(validationRequest.CreditCardNumber)
		isValid, _ := validator.Validate(cardNumber)

		data := ValidationResponse{IsValid: isValid}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
