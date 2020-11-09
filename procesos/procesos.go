package procesos

import (
	"fmt"
	"time"
)

// Procesos ...
func Procesos(id uint64, printChannel chan bool, stopProcess chan uint64) {
	i := uint64(0)
	for {
		select {
		case ProcessToStop := <-stopProcess:
			if id == ProcessToStop {
				// Terminamos el proceso.
				return
			}
		case <-printChannel:
			fmt.Printf("id %d: %d", id, i)
			fmt.Print("\n")
			i = i + 1
		default:
			i = i + 1
		}
		time.Sleep(time.Millisecond * 500)
	}
}

// Print ...
func Print(printChannel chan bool, deactivateFlag chan bool) {
	for {
		select {
		case <-deactivateFlag:
			for len(printChannel) > 0 {
				<-printChannel
				// Vaciamos el canal para que deje de imprimir.
			}
			return
			// Detenemos rutina y con ello la impresion.
		default:
			printChannel <- true
		}
	}
}

// Stop ...
func Stop(stopID uint64, procs []Proceso, stopProcess chan uint64) {
	for i := 0; i < len(procs); i++ {
		stopProcess <- stopID
	}
	// Para n = procesos en el array, mandamos n veces el id del proceso
	// a detener, para que en el peor de los casos, aun asi gracias al canal
	// se pueda identificar el proceso a detener, asi haya 10 o 1000 procesos
	// y asi el proceso a detener sea el primero en escuchar el canal (mejor caso),
	// o sea el ultimo en escuchar el canal (peor caso).
	return
}

// RemoveFromSlice ...
func RemoveFromSlice(id uint64, proc []Proceso) []Proceso {
	for s, element := range proc {
		if element.ID == id {
			return append(proc[:s], proc[s+1:]...)
		}
	}
	return proc
}

// Proceso ...
type Proceso struct {
	ID uint64
}
