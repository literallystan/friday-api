package main

import (
        "fmt"
        "log"
        "net/http"
        "github.com/gorilla/mux"
)

func main() {
    fmt.Println("hello world")
	router := mux.NewRouter()
    log.Fatal(http.ListenAndServe(":8000", router))
}
