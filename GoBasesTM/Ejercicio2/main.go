package main

import (
	"errors"
	"fmt"
)

func validarSalario(i int) (string, error) {
	if i < 150000 {
		return "", errors.New("error: el salario ingresado no alcanza el mínimo disponible")
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
