package main

import (
	"fmt"
	"os"
)

func LeerArchivo(nombre string) error {
	archivo, err := os.Open(nombre)
	if err != nil {
		fmt.Println("❌ Error al abrir el archivo:", err)
		return err
	}
	defer archivo.Close() // ¡Se ejecutará SIEMPRE al final!

	// Hacer cosas con el archivo...
	datos := make([]byte, 100)
	fmt.Println("Leyendo archivo:", nombre)
	_, err = archivo.Read(datos)
	if err != nil {
		return err // ¡archivo.Close() ya está asegurado!
	}

	return nil // ¡archivo.Close() se ejecuta aquí también!
}
func main() {
	LeerArchivo("README.md")
}
