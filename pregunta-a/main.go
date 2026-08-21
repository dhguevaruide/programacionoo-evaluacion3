package main

import (
	"fmt"
	"time"
)

func main() {
	// Creamos un rate limiter: 1 operación cada 500ms
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	// Simulamos 10 peticiones entrantes
	requests := make(chan int, 10)
	for i := 0; i < 10; i++ {
		requests <- i
	}
	close(requests)

	// Procesamos las peticiones con el rate limiter
	for req := range requests {
		<-ticker.C // ESPERAMOS el próximo TICK (oportunidad)
		fmt.Println("Procesando petición:", req, time.Now())
	}
}
