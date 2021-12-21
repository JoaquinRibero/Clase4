package main

import (
	"fmt"
)

func validarSalario(i int) (string, error) {
	if i < 150000 {
		err := fmt.Errorf("error: el mínimo imponible es de 150.000 y el salario ingresado es de: %d", i)
		return "", err
	} else {
		return "Debe pagar impuesto", nil
	}
}

func main() {
	salary := 145000
	s, err := validarSalario(salary)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(s)
	}
}
