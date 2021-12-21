package main

import (
	"errors"
	"fmt"
)

type MyError struct {
	msg string
}

func (e MyError) Error() string {
	return "error: el numero de meses trabajados no puede ser negativo"
}

func calcSalario(horas int, valor int) (float64, error) {
	salary := 0.0
	if horas < 80 {
		return 0.0, fmt.Errorf("error: el trabajador no puede haber trabajado menos de 80hs")
	} else {
		salary = float64(horas * valor)
		if salary > 150000 {
			salary = salary * 0.9
		}
		return salary, nil
	}
}

func calcAguinaldo(mSalario float64, mesesTrabajados int) (float64, error) {
	if mSalario < 0 || mesesTrabajados < 0 {
		return 0.0, MyError{}
	} else {
		aguinaldo := mSalario / 12 * float64(mesesTrabajados)
		return aguinaldo, nil
	}
}

func main() {

	s, err := calcSalario(78, 200)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("El salario es: ", s)
	}
	fmt.Println("--------------------")
	a, err := calcAguinaldo(150000, -1)
	if err != nil {
		e1 := fmt.Errorf("%w", err)
		fmt.Println(errors.Unwrap(e1))
	} else {
		fmt.Println("El aguinaldo es: ", a)
	}

}
