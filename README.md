Caesar Crypto
========================
[Caesar's Cipher implementation](https://en.wikipedia.org/wiki/Caesar_cipher) for Go

It is not considered cryptographicaly strong algorithm!
See also [ROT13](https://ru.wikipedia.org/wiki/ROT13) algorithm


Русский
========================
Реализация алгоритма [шифрования Цезаря](https://ru.wikipedia.org/wiki/%D0%A8%D0%B8%D1%84%D1%80_%D0%A6%D0%B5%D0%B7%D0%B0%D1%80%D1%8F) для языка Go.

Данный алгоритм не является надёжным!
См. также алгоритм [ROT13](https://ru.wikipedia.org/wiki/ROT13)

Badges / Значки
========================
[![Build Status](https://travis-ci.org/vodolaz095/vdlog.png?branch=master)](https://travis-ci.org/vodolaz095/caesar)
[![GoDoc](https://godoc.org/github.com/vodolaz095/vdlog?status.svg)](https://godoc.org/github.com/vodolaz095/caesar)
[![Go Report Card](https://goreportcard.com/badge/github.com/vodolaz095/caesar)](https://goreportcard.com/report/github.com/vodolaz095/caesar)


Example / Пример
========================


```go

package main

import (
	"fmt"

	"github.com/vodolaz095/caesar"
)

func main() {
	cryptor := caesar.NewRussian()
	fmt.Println(cryptor.Decrypt("ИЫЛРУДХРТ", 8)) // => Будильник
}

```