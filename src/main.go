package main

import (
	"log"
	"net/http"
	"gm"
)

func main() {

	router := gm.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}