package main

import (
	"fmt"
)

var valor int = 5 // Variable global

func main() {
	ch := make(chan int)

	// Enviar diferentes tipos de valores
	go func() {
		ch <- 42         // Enviar literal
		ch <- valor      // Enviar variable
		ch <- calcular() // Enviar resultado de función
		ch <- 10         // ← Enviar el valor 10 específicamente
	}()

	// Recibir 4 valores
	for i := 0; i < 4; i++ {
		fmt.Println(<-ch)
	}
}

func calcular() int {
	return 100
}
