package handler

import (
	"encoding/json"
	"net/http"

	"card-validator/internal/validator"
)

type request struct {
	CardNumber string `json:"card_number"`
}

type response struct {
	Valid    bool   `json:"valid"`
	CardType string `json:"card_type,omitempty"`
	Message  string `json:"message,omitempty"`
}

func ValidateCard(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var req request
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	valid, cardType := validator.Validate(req.CardNumber)

	res := response{
		Valid:    valid,
		CardType: cardType,
	}

	if !valid {
		res.Message = "Invalid card number"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
