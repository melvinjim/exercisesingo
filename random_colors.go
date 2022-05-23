package main
import (
	"fmt"
	"math/rand"
	"time"
)

var colors = [16]string{"rojo","verde","azul","salmon","gris","negro","veige","morado","blanco","turquesa","amarillo","neon","melocoton","diamante","ruby","escarlata"}

func main() {
	numberRandom()
}

func numberRandom(){
	rand.Seed(time.Now().UnixNano())
	var random = randomInt(0, 15)

	if (random % 2 == 1){
		numberRandom()
	} else {
		fmt.Println(colors[random])
	}
}

func randomInt(min, max int) int {
	var number = min + rand.Intn(max-min)
	return number
} 