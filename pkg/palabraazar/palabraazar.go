package palabraazar

import (
	"math/rand"
	"time"
)

// GetPalabraAzar devuelve una palabra al azar de una lista predefinida.
func GetPalabraAzar() string {
	palabras := []string{"manzana", "pera", "plátano", "uva", "naranja"}
	rand.Seed(time.Now().UnixNano())
	indice := rand.Intn(len(palabras))
	return palabras[indice]
}
