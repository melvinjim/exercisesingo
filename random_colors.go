package main
import (
	"fmt"
	"crypto/rand"
	"math/big"
)

var colors = [16]string{"rojo","verde","azul","salmon","gris","negro","veige","morado","blanco","turquesa","amarillo","neon","melocoton","diamante","ruby","escarlata"}

func main() {
	numberRandom()
}

func numberRandom(){
	var random = randomInt(0, 15)

	if (random % 2 == 1){
		numberRandom()
	} else {
		fmt.Println(colors[random])
	}
}

func randomInt(min, max int) int {
	nBig, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
    if err != nil {
        panic(err)
    }

    n := nBig.Int64()

	return int(n)
} 