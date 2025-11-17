package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"

	"github.com/atotto/clipboard"
	"github.com/fatih/color"
)

const Letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const Digits = "0123456789"
const Symbols = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"

func generatePassword(length int, digits bool, symbols bool) string {
	password := make([]rune, length)

	pool := Letters

	if digits {
		pool += Digits
	}
	if symbols {
		pool += Symbols
	}

	poolLength := big.NewInt(int64(len([]rune(pool))))
	poolRunes := []rune(pool)

	for i := range length {
		n, err := rand.Int(rand.Reader, poolLength)
		if err != nil {
			panic(err)
		}
		password[i] = poolRunes[n.Int64()]
	}
	return string(password)

}

func main() {
	length := flag.Int("length", 12, "Length of the password (default: 12)")
	digits := flag.Bool("digits", true, "Include digits in the password")
	symbols := flag.Bool("symbols", false, "Include symbols in the password")
	copy := flag.Bool("copy", true, "Enable or disable copying the password to the clipboard")

	flag.Parse()

	password := generatePassword(*length, *digits, *symbols)

	if *copy {
		clipboard.WriteAll(password)
		fmt.Print("Password copied! ")
	}

	color.Green("%s", password)
}
