package main

import (
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/api/pay", paymentHandler)
    log.Println("Server started on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func paymentHandler(w http.ResponseWriter, r *http.Request) {
    req, err := JSONBodyHandler(w, r)
    if err != nil {
        return
    }

    rate, err := GetExchangeRate(req.Currency)
    if err != nil {
        http.Error(w, "Failed to get exchange rate", http.StatusInternalServerError)
        return
    }

    convertedAmount := req.Amount * rate
    response := map[string]interface{}{
        "success":        true,
        "convertedAmount": convertedAmount,
        "currency":      req.Currency,
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}
