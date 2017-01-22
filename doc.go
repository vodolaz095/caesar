/*
Package caesar is a Golang implementation of Caesar's Cipher. This algorithm has some historical interest but is not considered cryptographicaly strong
It also implements ROT13 algorithm as subset of Caesar's Cipher with shift of 13.
Symbols not defined in Alphabet are passed as is.
See https://en.wikipedia.org/wiki/Caesar_cipher and https://en.wikipedia.org/wiki/ROT13

Русский: реализация алгоритма шифрования Цезаря для языка Go. Представляет исторический интерес, но не является криптографически стойким!
Также реализует алгоритм шифрования ROT13 как подвид алгоритма шифрования Цезаря со смещением 13 для английского языка.
Символы не заданные в алфавите передаются как есть.

См. https://ru.wikipedia.org/wiki/%D0%A8%D0%B8%D1%84%D1%80_%D0%A6%D0%B5%D0%B7%D0%B0%D1%80%D1%8F и https://ru.wikipedia.org/wiki/ROT13
*/
package caesar
