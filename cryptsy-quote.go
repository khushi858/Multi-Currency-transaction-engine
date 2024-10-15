package main

import (
    "net/http"
    "encoding/json"
)

type CryptsyResponse struct {
    Success int `json:"success"`
    Return  map[string]map[string]float64 `json:"return"`
}

func GetExchangeRate(currency string) (float64, error) {
    url := "https://api.cryptsy.com/api/marketdata"

    resp, err := http.Get(url)
    if err != nil {
        return 0, err
    }
    defer resp.Body.Close()

    var result CryptsyResponse
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return 0, err
    }

    if result.Success == 1 {
        return result.Return[currency]["lasttradeprice"], nil
    }

    return 0, nil
}
