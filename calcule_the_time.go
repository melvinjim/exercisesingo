package main 

import (
	"fmt"
)

func main(){
	operation(12, 7)
	operation(7, 15)
	operation(2, 39)
}

func operation(hours, minutes float64) float64{
	calculoM := (minutes / 60)
	total := float64(calculoM) + hours
	horas_totales := (total * 30)
	min_grados := (minutes * 6)
	totalResult := (min_grados - horas_totales)
	if (totalResult < 0){
		fmt.Println(int(-totalResult))
	} else {
		fmt.Println(int(totalResult))
	}
	
	return totalResult
}
