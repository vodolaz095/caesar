package caesar

import (
	"strings"
)

//Caesar is a struct, holding alphabet being used, alongside with Encrypt and Decrypt functions
type Caesar struct {
	Alphabet []string
}

//RussianAlphabet is a modern Russian language alphabet - with Ё
const RussianAlphabet = "АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ"

//EnglishAlphabet is a modern English language alphabet
const EnglishAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

//LatinAlphabet is a Classical Latin Alphabet
const LatinAlphabet = "ABCDEFGHIKLMNOPQRSTVXYZ"

//New makes interface for encrypting/decrypting any alphabet consisting from manifold of UTF8 characters
func New(alphabet string) Caesar {
	return Caesar{Alphabet: strings.Split(alphabet, "")}
}

//NewRussian makes interface for encrypting/decrypting  Russian alphabet
func NewRussian() Caesar {
	return New(RussianAlphabet)
}

//NewEnglish makes interface for encrypting/decrypting  English alphabet
func NewEnglish() Caesar {
	return New(EnglishAlphabet)
}

//NewLatin makes interface for encrypting/decrypting  Classic Latin alphabet
func NewLatin() Caesar {
	return New(LatinAlphabet)
}

func (c *Caesar) doShiftedAlphabed(shift int) (shiftedAlphabet []string) {
	shiftedAlphabet = make([]string, len(c.Alphabet))

	if shift <= 0 {
		for shift < len(c.Alphabet) {
			shift = len(c.Alphabet) + shift
		}
	}

	for shift > len(c.Alphabet) {
		for shift > len(c.Alphabet) {
			shift = shift - len(c.Alphabet)
		}
	}

	for ak := range c.Alphabet {
		if (ak + shift) < len(c.Alphabet) {
			shiftedAlphabet[ak] = c.Alphabet[ak+shift]
		} else {
			shiftedAlphabet[ak] = c.Alphabet[ak+shift-len(c.Alphabet)]
		}
	}
	return
}

func (c *Caesar) findInAlphabet(alphabet []string, input string) int {
	for k, v := range alphabet {
		if v == input {
			return k
		}
	}
	return -1
}

//Encrypt works for given string with shift provided
func (c *Caesar) Encrypt(input string, shift uint) (ret string) {
	d := int(shift)
	shiftedAlphabet := c.doShiftedAlphabed(d)
	for _, v := range strings.Split(input, "") {
		position := c.findInAlphabet(c.Alphabet, v)
		if position != -1 {
			ret = ret + shiftedAlphabet[position]
		} else {
			ret = ret + v
		}
	}
	return
}

//Decrypt works for given string with shift provided
func (c *Caesar) Decrypt(input string, shift uint) (ret string) {
	d := int(shift)
	shiftedAlphabet := c.doShiftedAlphabed(d)
	for _, v := range strings.Split(input, "") {
		position := c.findInAlphabet(shiftedAlphabet, v)
		if position != -1 {
			ret = ret + c.Alphabet[position]
		} else {
			ret = ret + v
		}
	}
	return
}
