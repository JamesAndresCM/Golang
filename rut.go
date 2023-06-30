// Usage go run rut.go 11111111
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const MOD_DV = 11

type Rut struct {
	Input int
	DV    interface{}
}

func (rut Rut) IsValid() bool {
	input := len(strconv.Itoa(rut.Input))
	if input >= 7 && input <= 8 {
		return true
	} else {
		return false
	}
}

func reverseArray(arr []string) []int {
	var intSlice []int
	for _, str := range arr {
		num, _ := strconv.Atoi(str)
		intSlice = append(intSlice, num)
	}

	for i, j := 0, len(intSlice)-1; i < j; i, j = i+1, j-1 {
		intSlice[i], intSlice[j] = intSlice[j], intSlice[i]
	}

	var reversedDV []int
	for _, num := range intSlice {
		reversedDV = append(reversedDV, num)
	}

	return reversedDV
}

func sumArray(arr []int) int {
	var i int = 0
	for _, value := range arr {
		i += value
	}

	return i
}

func (rut *Rut) formatDV(res int) {
	switch res {
	case 11:
		rut.DV = 0
	case 10:
		rut.DV = 'k'
	default:
		rut.DV = res
	}
}

func (rut Rut) CalcDV() Rut {
	dv := strconv.Itoa(rut.Input)
	splitDV := strings.Split(dv, "")
	reversedDV := reverseArray(splitDV)

	var i int = 1
	var j int = 1
	var totalSum []int

	for _, value := range reversedDV {
		i += 1
		if i < 8 {
			totalSum = append(totalSum, value*i)
		} else {
			j += 1
			totalSum = append(totalSum, value*j)
		}
	}
	sum := sumArray(totalSum)
	res := (MOD_DV - (sum % MOD_DV))

	rut.formatDV(res)
	return rut
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Debes ingresar un rut para calcular su Digito Verificador")
		return
	}

	input, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("El valor de Input debe ser un número entero válido")
		return
	}

	rut := Rut{Input: input}
	if !(rut.IsValid()) {
		fmt.Println("el rut ingresado no es valido")
		os.Exit(1)
	}
	rut = rut.CalcDV()
	fmt.Printf("Para el rut ingresado %d el digito verificador es %d\n", rut.Input, rut.DV)
}
