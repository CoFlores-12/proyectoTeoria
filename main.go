package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"priorityqueue/priorityqueue"
	"strconv"
	"time"
)

type Ticket struct {
	id        string
	arrival   string
	startTime string
	endTime   string
}

type Cajero struct {
	id       int
	busy     bool
	time1    time.Time
	timeBusy time.Duration
}

var queue *priorityqueue.PriorityQueue

var nCajeros int = 2
var cajeros = make([]*Cajero, nCajeros)

func main() {

	rand.Seed(time.Now().UnixNano())

	queue = priorityqueue.NewPriorityQueue(2)

	// iniciar los cajeros
	for i := 0; i < nCajeros; i++ {
		timeStart := time.Now()
		cajeros[i] = &Cajero{id: i + 1, busy: false, time1: timeStart, timeBusy: 0 * time.Second}
	}
	initServer()

}

// ############################## Server ##############################
func initServer() {
	mux := http.NewServeMux()

	go asignarTickets()
	mux.Handle("/", &homeHandler{})
	http.ListenAndServe(":8080", mux)

}

type homeHandler struct{}

func (h *homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path == "/NewTicket" {
		randId := rand.Intn(100)
		generarTicket(randId)
		w.Write([]byte("Hello world"))
		response := fmt.Sprintf("Ingreso el Ticket con ID: %d", randId)
		w.Write([]byte(response))
	} else {

		http.NotFound(w, r)
	}
}

//############################## Server ##############################

func generarTicket(i int) {
	randriority := rand.Intn(2) + 1

	// Obtener la hora actual
	currentTime := time.Now()

	// Formatear la hora actual en el formato "HH:MM:SS"
	timeString := fmt.Sprintf(currentTime.Format("15-04-05"))

	id := strconv.Itoa(i)

	min := 30  // 30 segundos
	max := 600 // 10  minutos
	// se calcula aleatoriamente el tiempo que va a tardar en aternderse dicho ticket
	duracion := strconv.Itoa(rand.Intn(max-min+1) + min)

	newTicket(randriority, id, timeString, duracion)
}

func asignarTickets() {
	//asignar tickets
	for true {

		ticket := queue.Pop()
		if ticket == nil {
			continue
		}
		// Convertir el tiempo
		numDuracion, err := strconv.Atoi(ticket.StartTime)
		if err != nil {
			// Manejar el error si la conversión falla
			fmt.Println("Error:", numDuracion)
			return
		}

		newTimeString, err := sumarSegundos(ticket.Arrival, numDuracion)
		if err != nil {
			fmt.Println("Error al sumar segundos:", err)
			return
		}

		fmt.Printf("ID: %s, Entro: %s, se va a tardar: %s, con prioridad: %d va a salir: %s\n", ticket.ID, ticket.Arrival, ticket.StartTime, ticket.Priority, newTimeString)
		fillLog(ticket.ID, 1, ticket.Priority, ticket.Arrival, newTimeString, ticket.Arrival)

	}
}

func parseTimeString(timeString string) (time.Time, error) {
	layout := "15-04-05"
	return time.Parse(layout, timeString)
}

func sumarSegundos(timeString string, secondsToAdd int) (string, error) {
	parsedTime, err := parseTimeString(timeString)
	if err != nil {
		return "", err
	}

	newTime := parsedTime.Add(time.Duration(secondsToAdd) * time.Second)
	return newTime.Format("15-04-05"), nil
}

// funcion para crear unu ticket
func newTicket(priority int, id string, arrival string, startTime string) *string {
	ticket := &Ticket{
		id:        id,
		arrival:   arrival,
		startTime: startTime,
	}

	if priority == 1 {
		id = "T" + id
	} else {
		id = "R" + id
	}

	queue.Push(ticket, priority, id, arrival, startTime)
	return &id
}

// funcion para llenar el reports.py
func fillLog(ticketID string, waitTime int, priority int, startTime string, endTime string, serveTime string) {
	// llenar el log
	// abrir el archivo de registro
	logFile, err := os.OpenFile("registro.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		fmt.Printf("Error abriendo el archivo de registro: %v\n", err)
		return
	}
	defer logFile.Close()

	// ingresar datos al arcchivo
	logMessage := fmt.Sprintf("id:%s,wd:%d,priority:%d,t1:%s,t2:%s,t3:%s\n",
		ticketID, waitTime, priority, startTime, endTime, serveTime)
	logFile.WriteString(logMessage)

}
