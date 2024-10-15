package main

import (
    "encoding/json"
    "net/http"
)

func paymentHandler(w http.ResponseWriter, r *http.Request) {
    var request map[string]interface{}
    json.NewDecoder(r.Body).Decode(&request)
    // Implement payment logic here
    response := map[string]interface{}{
        "id":      11022047,
        "address": "1FfWNa82MyfZqVhQwW7yRM7ZTxzS1V9hx8",
        "currency": "BTC",
        "amount":   3050,
        "fees":     50,
        "timeout":  10000,
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}
