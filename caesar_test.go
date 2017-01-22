package caesar

import (
	"fmt"
	"strings"
	"testing"
)

type testVectorAlphabet struct {
	Shift  int
	New    string
	Result string
}

func TestItShiftsAlphabed(t *testing.T) {
	tst := []testVectorAlphabet{
		{1, "ABCD", "BCDA"},
		{2, "ABCD", "CDAB"},
		{3, "ABCD", "DABC"},
		{4, "ABCD", "ABCD"},
		{-1, "ABCD", "DABC"},

		{1, "АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ", "БВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯА"},
		{2, "АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ", "ВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯАБ"},
		{3, "АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ", "ГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯАБВ"},
		{4, "АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ", "ДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯАБВГ"},
		{5, "АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ", "ЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯАБВГД"},
	}

	for _, v := range tst {
		simple := New(v.New)
		ret := strings.Join(simple.doShiftedAlphabed(v.Shift), "")
		if ret != v.Result {
			t.Errorf("error shifting alphabet %s for %v. Got %s instead of %s", v.New, v.Shift, ret, v.Result)
		}
	}

}

func TestFindInAlphabet(t *testing.T) {
	simple := New("ABCD")
	ret := simple.findInAlphabet(simple.Alphabet, "A")
	if ret != 0 {
		t.Errorf("error finding in alphabet %s : %s has position of %v instead of %v", strings.Join(simple.Alphabet, ""), "A", ret, 0)
	}

	ret = simple.findInAlphabet(simple.Alphabet, "G")
	if ret != -1 {
		t.Errorf("error finding in alphabet %s : %s has position of %v instead of %v", strings.Join(simple.Alphabet, ""), "G", ret, -1)
	}

	ret = simple.findInAlphabet(simple.Alphabet, "B")
	if ret != 1 {
		t.Errorf("error finding in alphabet %s : %s has position of %v instead of %v", strings.Join(simple.Alphabet, ""), "B", ret, 1)
	}
}

func TestEncrypt(t *testing.T) {
	simple := New("ABCDEF")
	encrypted := simple.Encrypt("AA", 3)
	if encrypted != "DD" {
		t.Errorf("wrong encrypted string - got %s instead of DD", encrypted)
	}
}

func TestDecrypt(t *testing.T) {
	simple := New("ABCDEF")
	decrypted := simple.Decrypt("DD", 3)
	if decrypted != "AA" {
		t.Errorf("wrong decrypted string - got %s instead of AA", decrypted)
	}
}

type testVector struct {
	Iteration uint
	Decrypted string
	Encrypted string
}

func TestRot13(t *testing.T) {
	rot13 := NewEnglish()
	if rot13.Encrypt("HELLO", 13) != "URYYB" {
		t.Error("rot13 is broken")
	}
}

func TestAll(t *testing.T) {
	iterations := []testVector{
		{0, "ИЫЛРУДХРТ", "ИЫЛРУДХРТ"},
		{1, "ЗЪКПТГФПС", "ЙЬМСФЕЦСУ"},
		{2, "ЖЩЙОСВУОР", "КЭНТХЁЧТФ"},
		{3, "ЁШИНРБТНП", "ЛЮОУЦЖШУХ"},
		{4, "ЕЧЗМПАСМО", "МЯПФЧЗЩФЦ"},
		{5, "ДЦЖЛОЯРЛН", "НАРХШИЪХЧ"},
		{6, "ГХЁКНЮПКМ", "ОБСЦЩЙЫЦШ"},
		{7, "ВФЕЙМЭОЙЛ", "ПВТЧЪКЬЧЩ"},
		{8, "БУДИЛЬНИК", "РГУШЫЛЭШЪ"},
		{9, "АТГЗКЫМЗЙ", "СДФЩЬМЮЩЫ"},
	}

	hard := New("АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ")

	for _, v := range iterations {
		if v.Iteration == 0 {
			continue
		}

		enc := hard.Encrypt("ИЫЛРУДХРТ", v.Iteration)
		if enc != v.Encrypted {
			t.Errorf("error encrypting got >%s< instead of >%s< for %v iteration %t", enc, v.Encrypted, uint(v.Iteration), enc == v.Encrypted)
		}
		dec := hard.Decrypt("ИЫЛРУДХРТ", v.Iteration)
		if dec != v.Decrypted {
			t.Errorf("error decrypting got >%s< instead of >%s< for %v iteration %t", dec, v.Decrypted, uint(v.Iteration), dec == v.Decrypted)
		}
	}
}

func TestRespectUnknownSymbols(t *testing.T) {
	rot13 := NewEnglish()
	if rot13.Encrypt("HELLO!", 13) != "URYYB!" {
		t.Error("unknown symbols mangled")
	}
}

func ExampleNewRussian() {
	cryptor := NewRussian()
	fmt.Println(cryptor.Decrypt("ИЫЛРУДХРТ", 8)) // => Будильник
}

func Example() {
	//for rot13 algorithm
	cryptor := NewEnglish()
	fmt.Println(cryptor.Decrypt("URYYB", 13)) // => HELLO
	fmt.Println(cryptor.Encrypt("HELLO", 13)) // => URYYB

	//for Russian language
	cryptorR := NewRussian()
	fmt.Println(cryptorR.Decrypt("ИЫЛРУДХРТ", 8)) // => БУДИЛЬНИК
	fmt.Println(cryptorR.Encrypt("БУДИЛЬНИК", 8)) // => ИЫЛРУДХРТ

}
