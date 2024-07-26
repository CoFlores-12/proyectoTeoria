package main

import (
	"fmt"
	"math/rand"
	"os"
	"priorityqueue/priorityqueue"
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

func main() {
	nCajeros := 2
	rand.Seed(time.Now().UnixNano())

	queue = priorityqueue.NewPriorityQueue(2)
	cajeros := make([]*Cajero, nCajeros)

	// iniciar los cajeros
	for i := 0; i < nCajeros; i++ {
		timeStart := time.Now()
		cajeros[i] = &Cajero{id: i + 1, busy: false, time1: timeStart, timeBusy: 0 * time.Second}
	}

	//generar tickets
	for i := 0; i < 2; i++ {
		randriority := rand.Intn(3)
		hour := rand.Intn(24)
		minute := rand.Intn(60)
		second := rand.Intn(60)
		timeString := fmt.Sprintf("%02d:%02d:%02d", hour, minute, second)
		newTicket(randriority, timeString)
	}

	//asignar tickets
	for i := 0; i < nCajeros; i++ {
		tiempoActual := time.Now()
		tiempoTranscurrido := tiempoActual.Sub(cajeros[i].time1)
		if tiempoTranscurrido > cajeros[i].timeBusy {
			continue
		}
		ticket := queue.Pop()
		if ticket != nil {
			fmt.Println("Ticket:", ticket)
			fmt.Println("Data:", ticket.Data)
		}
	}
}

// funcion para crear unu ticket
func newTicket(priority int, timeA string) *string {
	//TODO: generar ID, T1 terceraEdad, R1 Regular
	var id = ""
	ticket := &Ticket{
		id:      id,
		arrival: timeA,
	}
	queue.Push(ticket, priority)
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
