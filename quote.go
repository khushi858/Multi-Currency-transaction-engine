package main

import (
    "encoding/json"
    "net/http"
)

func quoteHandler(w http.ResponseWriter, r *http.Request) {
    var request map[string]interface{}
    json.NewDecoder(r.Body).Decode(&request)
    // Implement currency conversion logic here
    response := map[string]interface{}{
        "BTC": 3172000,
        "ETH": 2370000,
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}
