package main

import (
	"go-axesthump-adventure/internal/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	handler, err := handlers.NewAppHandler()
	if err != nil {
		exit(err.Error())
	}
	err = http.ListenAndServe(":8080", handler.Router)
	if err != nil {
		exit(err.Error())
	}
}

func exit(msg string) {
	log.Println(msg)
	os.Exit(1)
}
