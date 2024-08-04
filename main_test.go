package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"testing"
)

func TestNewTicketEndpoint(t *testing.T) {
	numRequests := 10 // Número de requests que quieres hacer
	totalIDs := 0

	for i := 0; i < numRequests; i++ {
		// Generar un número aleatorio entre 1 y 2 para la prioridad
		prioridad := rand.Intn(2) + 1

		// Crear el body de la solicitud con el parámetro "prioridad"
		body := []byte(fmt.Sprintf(`{"priority": %d}`, prioridad))

		resp, err := http.Post("http://localhost:8080/NewTicket", "application/json", bytes.NewBuffer(body))
		if err != nil {
			log.Fatalf("Failed to make request: %v", err)
		}
		defer resp.Body.Close()

		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("Failed to read response body: %v", err)
		}

		var response struct {
			Message string `json:"message"`
			ID      int    `json:"id"`
		}

		if err := json.Unmarshal(bodyBytes, &response); err != nil {
			log.Fatalf("Failed to unmarshal response body: %v", err)
		}

		// Verificar que se recibió un ID válido
		if response.ID > 0 {
			totalIDs++
		} else {
			t.Errorf("No se recibió un ID válido en la respuesta: %s", string(bodyBytes))
		}
	}

	expectedTotalIDs := numRequests // El total de IDs esperado es igual al número de requests
	if totalIDs != expectedTotalIDs {
		t.Errorf("Total de IDs recibidos no es igual al total de requests. Total IDs: %d, Total Requests: %d", totalIDs, expectedTotalIDs)
	}
}
