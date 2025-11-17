package main

import (
	"flag"
	"fmt"
	"math/rand"

	"github.com/atotto/clipboard"
	"github.com/fatih/color"
)

const Letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const Digits = "0123456789"
const Symbols = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"

func generatePassword(length int, digits bool, symbols bool) string {
	password := make([]rune, length)

	for i := range length {
		options := Letters

		if digits {
			options += Digits
		}
		if symbols {
			options += Symbols
		}
		password[i] = rune(options[rand.Intn(len(options))])
	}
	return string(password)

}

func main() {
	length := flag.Int("length", 12, "Length of the password (default: 12)")
	digits := flag.Bool("digits", true, "Include digits in the password")
	symbols := flag.Bool("symbols", false, "Include symbols in the password")
	flag.Parse()

	password := generatePassword(*length, *digits, *symbols)

	clipboard.WriteAll(password)

	fmt.Print("Password copied! ")
	color.Green("%s", password)
}
