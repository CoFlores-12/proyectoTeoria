package main

import (
	// "testing"

	"fmt"
	"math/rand"
	"os"
	priorityqueue "priorityqueue/PriorityQueue"
	"time"
)

// func TestFunctions(t *testing.T) {

// }

// func TestLen(t *testing.T) {

// 	ln := 0
// 	if ln != 8 {
// 		t.Errorf("Esperado: %v / Obtenido: %v", 8, ln)
// 	}
// }

type Ticket struct {
	id        string
	priority  int
	arrival   string
	startTime string
	endTime   string
}

type Cajero struct {
	id   int
	busy bool
}

var queue *priorityqueue.PriorityQueue

func main() {
	nCajeros := 2
	nTickets := 10

	queue = priorityqueue.NewPriorityQueue(2)
	cajeros := make([]*Cajero, nCajeros)

	// iniciar los cajeros
	for i := 0; i < nCajeros; i++ {
		cajeros[i] = &Cajero{id: i + 1, busy: false}
	}

	// Generar y agregar tickets a la cola
	for i := 0; i <= nTickets; i++ {
		isThirdAge := rand.Intn(2) == 0
		newTicket(i, isThirdAge)
	}

	// Atender tickets
	for i := 0; queue.GetLenElements() > 0; i++ {

		cajero := rand.Intn(3)

		item := queue.Pop()
		serveTime := time.Duration(rand.Intn(5)+1) * time.Second
		endTime1 := time.Now().Add(time.Duration(serveTime) * time.Second).Format("15-04-05")

		fmt.Printf(" atendido por: %d termino a las: %s\n", cajero, endTime1, item)

		// switch servidor {
		// case 0:

		// 	break
		// }

		// 	if banderaCajero1 {
		// 		cajeros[t].busy = true
		// 		item := queue.Pop()

		// 	}

		// if cajeros[t].busy == false {
		// 	cajeros[t].busy = true
		// 	item := queue.Pop()

		// 	fmt.Printf("\n", item)

		// 	// print("cajero", t)
		// 	cajeros[t].busy = false
		// 	serveTime := time.Duration(rand.Intn(5)+1) * time.Second
		// 	time.Sleep(serveTime)
		// 	endTime := time.Now().Format("15-04-05")
		// 	print(endTime)
		// }

		// time1 := time.Now().Format("15:04:05")

		// switch true {
		// case cajeros[0].busy == false:
		// 	cajeros[0].busy = true
		// 	item1 := queue.Pop()

		// 	print("cajero", 0)
		// 	tiempo1 := time1

		// 	serveTime := time.Duration(rand.Intn(5)+1) * time.Second
		// 	// time.Sleep(serveTime)

		// 	endTime1 := time.Now().Add(time.Duration(serveTime) * time.Second).Format("15:04:05")

		// 	if endTime1 == tiempo1 {
		// 		cajeros[0].busy = false
		// 		fmt.Printf("\n", item1)
		// 		print(endTime1)
		// 	}

		// 	break

		// }

		// // prueba
		// ticketID := "T1"
		// waitTime := 5
		// priority := 1
		// startTime := time.Now().Format("15-04-05")
		// endTime := time.Now().Add(2 * time.Minute).Format("15-04-05")
		// serveTime := time.Now().Add(6 * time.Minute).Format("15-04-05")

		// //llenado del .log
		// fillLog(ticketID, waitTime, priority, startTime, endTime, serveTime)

	}

}

// funcion para crear unu boleto
func newTicket(id int, isThirdAge bool) {
	t := time.Now()
	time := t.Format("15-04-05")
	prio := 1

	if isThirdAge == true {
		prio = 1
	} else {
		prio = 2
	}

	ticket := &Ticket{
		id:       fmt.Sprintf("T%d", id),
		priority: prio,
		arrival:  time,
	}
	if isThirdAge {
		ticket.priority = priorityqueue.ThirdAgePriority
	}
	queue.Push(ticket, isThirdAge)
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
