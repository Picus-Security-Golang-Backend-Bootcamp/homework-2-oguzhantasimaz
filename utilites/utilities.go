package utilites

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

//random string generator with cryptorand
func RandomString(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		value, err := rand.Int(rand.Reader, big.NewInt(int64(len(letterRunes))))
		if err != nil {
			fmt.Println(err)
		}
		b[i] = letterRunes[value.Int64()]
	}
	return string(b)
}

//random int generator with cryptorand
func RandomInt(max int64) int {
	value, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		fmt.Print(err)
	}
	return int(value.Int64())
}
