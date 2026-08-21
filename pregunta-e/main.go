package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	count := 0
	for {
		select {
		case <-ticker.C:
			count++
			fmt.Println("Tick:", count)
			if count >= 3 {
				return // ← ¿Qué pasa aquí?
			}
		}
	}
}
