package main

import (
    "encoding/json"
    "net/http"
)

type PaymentRequest struct {
    Amount   float64 `json:"amount"`
    Currency string  `json:"currency"`
}

func JSONBodyHandler(w http.ResponseWriter, r *http.Request) (*PaymentRequest, error) {
    var req PaymentRequest
    decoder := json.NewDecoder(r.Body)
    err := decoder.Decode(&req)
    if err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return nil, err
    }
    return &req, nil
}
