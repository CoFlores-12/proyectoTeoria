package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"priorityqueue/priorityqueue"
	"strconv"
	"time"
)

type Ticket struct {
	id      string
	arrival string
}

type Cajero struct {
	id       int
	busy     bool
	time1    string
	timeBusy int
}

var queue *priorityqueue.PriorityQueue

var nCajeros int = 2
var ID int = 1
var cajeros = make([]*Cajero, nCajeros)

func main() {

	rand.Seed(time.Now().UnixNano())

	queue = priorityqueue.NewPriorityQueue(2)

	// iniciar los cajeros
	for i := 0; i < nCajeros; i++ {
		timeStart := time.Now()
		cajeros[i] = &Cajero{id: i + 1, busy: false, time1: timeStart.Format("15-04-05"), timeBusy: 0}
	}
	initServer()

}

// ############################## Server ##############################

type homeHandler struct{}

type TicketRequest struct {
	Priority int `json:"priority"`
}

func (h *homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost && r.URL.Path == "/NewTicket" {
		var req TicketRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if req.Priority < 1 || req.Priority > 2 {
			http.Error(w, "Invalid priority", http.StatusBadRequest)
			return
		}

		var id string

		if req.Priority == 1 {
			id = "T" + strconv.Itoa(ID)
		} else {
			id = "R" + strconv.Itoa(ID)
		}

		generarTicket(id, req.Priority)
		response := fmt.Sprintf("Ingreso el Ticket con ID: %s y Prioridad: %d", id, req.Priority)
		ID++
		w.Write([]byte(response))
	} else {
		http.NotFound(w, r)
	}
}

func initServer() {
	fmt.Println("Servidor activo en http://localhost:8080")
	go asignarTickets()
	mux := http.NewServeMux()
	mux.Handle("/", &homeHandler{})
	http.ListenAndServe(":8080", mux)
}

//############################## Server ##############################

func generarTicket(id string, priority int) {
	// Obtener la hora actual
	currentTime := time.Now()

	// Formatear la hora actual en el formato "HH:MM:SS"
	timeString := fmt.Sprintf(currentTime.Format("15-04-05"))

	newTicket(priority, id, timeString)
}

func asignarTickets() {
	// llenar el log
	// abrir el archivo de registro
	logFile, err := os.OpenFile("reports.log", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)

	if err != nil {
		fmt.Printf("Error abriendo el archivo de registro: %v\n", err)
		return
	}
	defer logFile.Close()

	//asignar tickets
	for true {

		currentTime := time.Now()
		// Formatear la hora actual en el formato "HH-MM-SS"
		timeString := fmt.Sprintf(currentTime.Format("15-04-05"))

		if restarSegundos(timeString, cajeros[0].time1) >= cajeros[0].timeBusy {
			ticket := queue.Pop()
			if ticket == nil {
				continue
			}

			min := 3  // 30 segundos
			max := 20 // 10  minutos
			// se calcula aleatoriamente el tiempo que va a tardar en aternderse dicho ticket
			duracion := rand.Intn(max-min+1) + min

			cajeros[0].timeBusy = duracion

			cajeros[0].time1 = timeString

			fmt.Printf("ID: %s,lo atendio 1, con prioridad: %d, Entro: %s, se va a tardar: %d, se atendia a las: %s va a salir: %s\n", ticket.ID, ticket.Priority, ticket.Arrival, duracion, timeString, sumarSegundos(timeString, duracion))

			// ingresar datos al arcchivo
			logMessage := fmt.Sprintf("id:%s,wd:%d,priority:%d,t1:%s,t2:%s,t3:%s\n", ticket.ID, 1, ticket.Priority, ticket.Arrival, timeString, sumarSegundos(timeString, duracion))
			logFile.WriteString(logMessage)

		}

		if restarSegundos(timeString, cajeros[1].time1) >= cajeros[1].timeBusy {
			ticket := queue.Pop()
			if ticket == nil {
				continue
			}

			min := 30  // 30 segundos
			max := 600 // 10  minutos = 600 segundos
			// se calcula aleatoriamente el tiempo que va a tardar en aternderse dicho ticket
			duracion := rand.Intn(max-min+1) + min

			cajeros[1].timeBusy = duracion

			cajeros[1].time1 = timeString

			fmt.Printf("ID: %s,lo atendio 2, con prioridad: %d, Entro: %s, se va a tardar: %d, se atendia a las: %s va a salir: %s\n", ticket.ID, ticket.Priority, ticket.Arrival, duracion, timeString, sumarSegundos(timeString, duracion))
			logMessage := fmt.Sprintf("id:%s,wd:%d,priority:%d,t1:%s,t2:%s,t3:%s\n", ticket.ID, 2, ticket.Priority, ticket.Arrival, timeString, sumarSegundos(timeString, duracion))
			logFile.WriteString(logMessage)
		}

	}
}

func restarSegundos(hora1, hora2 string) int {
	t1, err := time.Parse("15-04-05", hora1)
	if err != nil {
		return 0
	}

	t2, err := time.Parse("15-04-05", hora2)
	if err != nil {
		return 0
	}

	diferencia := t1.Sub(t2)
	return int(diferencia.Seconds())
}

func parseTimeString(timeString string) (time.Time, error) {
	layout := "15-04-05"
	return time.Parse(layout, timeString)
}

func sumarSegundos(timeString string, secondsToAdd int) string {
	parsedTime, err := parseTimeString(timeString)
	if err != nil {
		return ""
	}

	newTime := parsedTime.Add(time.Duration(secondsToAdd) * time.Second)
	return newTime.Format("15-04-05")
}

// funcion para crear unu ticket
func newTicket(priority int, id string, arrival string) *string {
	ticket := &Ticket{
		id:      id,
		arrival: arrival,
	}

	queue.Push(ticket, priority, id, arrival)
	return &id
}
