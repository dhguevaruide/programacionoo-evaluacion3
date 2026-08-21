package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	// Lanzamos 3 goroutines
	for i := 1; i <= 3; i++ {
		wg.Add(1) // Incrementar contador antes de lanzar

		go func(id int) {
			defer wg.Done() // Asegurar que se decrementa
			fmt.Printf("Goroutine %d: trabajando...\n", id)
			time.Sleep(time.Duration(id) * time.Second)
			fmt.Printf("Goroutine %d: terminé\n", id)
		}(i)
	}

	fmt.Println("Esperando que todas terminen...")
	wg.Wait() // ← BLOQUEA hasta que contador = 0
	fmt.Println("¡Todas las goroutines terminaron!")
}
