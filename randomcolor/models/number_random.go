package models

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

var colors = [16]string{"rojo", "verde", "azul", "salmon", "gris", "negro", "veige", "morado", "blanco", "turquesa", "amarillo", "neon", "melocoton", "diamante", "ruby", "escarlata"}

func ColorRandom() string {
	var random = RandomInt(0, 15)

	if random%2 == 1 {
		ColorRandom()
	}
	return colors[random]
}

func RandomInt(min, max int) int {
	nBig, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		panic(err)
	}

	n := nBig.Int64()

	fmt.Println("n=", int(n))
	return int(n)
}
