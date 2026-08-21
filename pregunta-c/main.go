package main

import (
	"fmt"
	"os"
)

func main() {
	// 1. Abrir un archivo (puede fallar)
	archivo, err := os.Open("archivo_inexistente.txt")
	if err != nil {
		fmt.Println("❌ Error:", err)
		return // Salimos de la función
	}
	defer archivo.Close()

	// 2. Leer el archivo (puede fallar)
	datos := make([]byte, 100)
	_, err = archivo.Read(datos)
	if err != nil {
		fmt.Println("❌ Error al leer:", err)
		return
	}

	// 3. Si llegamos aquí, todo bien
	fmt.Println("✅ Archivo leído correctamente:", string(datos))
}
