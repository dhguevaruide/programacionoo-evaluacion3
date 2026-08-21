package main

import (
	"fmt"
)

func main() {
	// Creamos un canal
	ch := make(chan string)

	// Goroutine que envía mensajes
	go func() {
		ch <- "Mensaje 1"
		ch <- "Mensaje 2"
		ch <- "Mensaje 3"
		close(ch) // ¡IMPORTANTE! Cerramos el canal
	}()

	// range recibe TODOS los valores hasta que el canal se cierra
	for msg := range ch {
		fmt.Println("Recibido:", msg)
	}

	fmt.Println("Canal cerrado, terminamos")
}
