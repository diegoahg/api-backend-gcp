package api

import (
	"api-backend-gcp/pkg/palabraazar"
	"encoding/json"
	"net/http"
)

// API es una estructura para gestionar las rutas y controladores de la API.
type API struct{}

// NewAPI crea una nueva instancia de la API.
func NewAPI() *API {
	return &API{}
}

// HandlePalabraAzar maneja las solicitudes a la ruta /palabra-azar.
func (a *API) HandlePalabraAzar(w http.ResponseWriter, r *http.Request) {
	palabra := palabraazar.GetPalabraAzar()

	// Responder con la palabra en formato JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"palabra": palabra})
}

// ServeHTTP implementa el método ServeHTTP de la interfaz http.Handler.
func (a *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Llama a HandlePalabraAzar para manejar la solicitud
	a.HandlePalabraAzar(w, r)
}
