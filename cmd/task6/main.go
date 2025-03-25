package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func genPsw(length int) (string, error) {

	charSet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	var buf strings.Builder
	for i := 0; i < length; i++ {
		if err := buf.WriteByte(charSet[rand.Intn(len(charSet))]); err != nil {
			return "", err
		}
	}
	return buf.String(), nil
}

func main() {
	psw, err := genPsw(10)
	fmt.Println(psw, err)

}
