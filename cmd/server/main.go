package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	log.Fatal(http.ListenAndServe(":"+os.Getenv("SERVER_PORT"), nil))
}
