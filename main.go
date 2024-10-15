package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/quote/", quoteHandler).Methods("POST")
    router.HandleFunc("/payment/", paymentHandler).Methods("POST")
    router.HandleFunc("/status/", statusHandler).Methods("POST")

    log.Fatal(http.ListenAndServe(":8000", router))
}
