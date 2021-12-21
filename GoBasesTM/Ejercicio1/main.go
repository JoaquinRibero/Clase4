package main

import (
	"fmt"
)

type MyError struct{}

type error interface {
	Error() string
}

func (e MyError) Error() string {
	return "error: el salario ingresado no alcanza el mínimo disponible"
}

func validarSalario(i int) (string, error) {
	if i < 150000 {
		return "", MyError{}
	} else {
		return "Debe pagar impuesto", nil
	}
}

func main() {
	salary := 152000
	s, err := validarSalario(salary)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)

}
