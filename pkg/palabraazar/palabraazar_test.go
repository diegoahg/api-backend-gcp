package palabraazar

import (
	"testing"
)

func TestGetPalabraAzar(t *testing.T) {
	// Prueba que GetPalabraAzar devuelva una palabra válida
	palabra := GetPalabraAzar()
	if palabra == "" {
		t.Error("Se esperaba una palabra, pero se obtuvo una cadena vacía.")
	}
}
