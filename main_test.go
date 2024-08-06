package main

import (
	"bytes"
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

		// Crear el body de la solicitud con el parámetro "priority"
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

		responseBody := string(bodyBytes)
		fmt.Printf("Response Body: %s\n", responseBody)

		// Incrementar el total de IDs recibidos
		totalIDs++
	}

	expectedTotalIDs := numRequests // El total de IDs esperado es igual al número de requests
	if totalIDs != expectedTotalIDs {
		t.Errorf("Total de IDs recibidos no es igual al total de requests. Total IDs: %d, Total Requests: %d", totalIDs, expectedTotalIDs)
	}
}
