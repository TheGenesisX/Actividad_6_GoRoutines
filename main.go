package main

import (
	"fmt"

	"./procesos"
)

func main() {
	idProc := uint64(0)
	var opc, stop uint64
	var input string
	var procs []procesos.Proceso
	printChannel := make(chan bool)
	deactivateFlag := make(chan bool)
	stopProcess := make(chan uint64)

	for {
		fmt.Println("----------Administrador de procesos----------")
		fmt.Println("1) Agregar proceso")
		fmt.Println("2) Mostrar procesos")
		fmt.Println("3) Eliminar un proceso")
		fmt.Println("0) Salir")
		fmt.Print("Opcion: ")
		fmt.Scanln(&opc)

		if opc == 0 {
			for _, element := range procs {
				go procesos.Stop(element.ID, procs, stopProcess)
			}
			procs = nil
			fmt.Println("Saliendo del administrador.")
			break
		}

		switch opc {
		case 1:
			proc := new(procesos.Proceso)
			proc.ID = idProc
			idProc++
			procs = append(procs, *proc)

			go procesos.Procesos(proc.ID, printChannel, stopProcess)

		case 2:
			go procesos.Print(printChannel, deactivateFlag)

			fmt.Scanln(&input)
			// Hacemos que la concurrencia nos espere.

			deactivateFlag <- true
		case 3:
			fmt.Println("Proceso a detener: ")
			fmt.Scanln(&stop)
			go procesos.Stop(stop, procs, stopProcess)
			fmt.Println("Proceso detenido exitosamente")

			procs = procesos.RemoveFromSlice(stop, procs)

			fmt.Scanln(&input)
			// Hacemos que la concurrencia nos espere.
		}
	}
}
