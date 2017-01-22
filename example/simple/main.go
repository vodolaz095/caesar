package main

import (
	"fmt"

	"github.com/vodolaz095/caesar"
)

func main() {
	cryptor := caesar.NewRussian()
	fmt.Println(cryptor.Decrypt("ИЫЛРУДХРТ", 8)) // => Будильник
}
