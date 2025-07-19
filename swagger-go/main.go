package main

import (
	"log"
	"net/http"
	"swagger-go/router"
)

func main() {
	router.SetupRoutes()
	log.Println("Servidor iniciado en http://localhost:3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}
