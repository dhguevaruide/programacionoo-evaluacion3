package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	go func() {
		ch <- 1
		ch <- 2
		close(ch)
	}()

	// Recibir con verificación
	valor, ok := <-ch // ok = true si el canal está abierto
	if ok {
		fmt.Println("Recibido:", valor)
	}

	// O con range (automático)
	for valor := range ch { // ← range usa <-ch internamente
		fmt.Println("Recibido:", valor)
	}
}
