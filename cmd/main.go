package main

import (
	"api-backend-gcp/api"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Inicializar el servidor HTTP
	server := http.Server{
		Addr:    ":8080",
		Handler: api.NewAPI(),
	}

	// Iniciar el servidor
	fmt.Printf("Server Listening in port%s ...", server.Addr)
	log.Fatal(server.ListenAndServe())
}
