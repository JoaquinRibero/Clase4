package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.ReadFile(",/no-file.txt")
	if err != nil {
		panic("El archivo indicado no fue encontrado o está dañado")
	}

	fmt.Println("Fin del programa")
}
