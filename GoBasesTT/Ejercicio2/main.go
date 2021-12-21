package main

import (
	"errors"
	"fmt"
	"os"
)

func legajo(i int) int {
	i++
	return i
}

func readTxt(filename string) string {

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	_, err := os.ReadFile(filename)
	if err != nil {
		panic("error: el archivo indicado no fue encontrado o está dañado")
	} else {
		return ""
	}
}

func registrarCliente(legajo int, Nombre string, dni int, telefono int, domicilio string) (bool, error) {
	s := readTxt("./customers.txt")
	fmt.Println(s)

	if legajo == 0 || Nombre == "" || domicilio == "" || dni == 0 || telefono == 0 {
		return false, errors.New("Los datos del cliente no pueden estar vacíos o ser 0")
	} else {
		return true, nil
	}

}

func main() {

	i := 0 //-1 para generar el panic
	x := legajo(i)
	if x == 0 {
		panic("Hubo un error al generar el número de legajo.")
	}

	//la func registrar cliente llama a la que intenta leer el archivo inexistente
	_, err := registrarCliente(x, "Joaquin Ribero", 0, 3492123456, "Calle Falsa 123")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Fin de la ejecución")
	fmt.Println("Se detectaron varios errores en tiempo de ejecución.")
	fmt.Println("No han quedado archivos abiertos")

}
