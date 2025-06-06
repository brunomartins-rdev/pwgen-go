package generator

import (
	"math/rand"
	"time"

	"github.com/brunomartins-rdev/pwgen-go/internal/utils"
)

// TODO: Implement NumberOfGens
type PasswordOptions struct {
	Lenght           int
	IncludeLowercase bool
	IncludeUppercase bool
	IncludeNumbers   bool
	IncludeSymbols   bool
}

const lowercase = "abcdefghijklmnopqrstuvwxyz"
const uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const numbers = "1234567890"
const symbols = "!\"#$%&'()*+,-./:;<=>?@[\\]^_{|}~"

func GeneratePassword(options PasswordOptions) string {
	requiredCharacters := []byte{}
	allAllowedCharacters := ""

	if options.IncludeLowercase{
		requiredCharacters = append(
			requiredCharacters, 
			lowercase[rand.Intn(len(lowercase))]
		)

		allAllowedCharacters += lowercase
	}

	if options.IncludeUppercase {
		requiredCharacters = append(
			requiredCharacters, 
			uppercase[rand.Intn(len(uppercase))]
		)

		allAllowedCharacters += uppercase
	}

	if options.IncludeNumbers{
		requiredCharacters = append(
			requiredCharacters, 
			numbers[rand.Intn(len(numbers))]
		)

		allAllowedCharacters += numbers
	}

	if options.IncludeSymbols{
		requiredCharacters = append(
			requiredCharacters, 
			symbols[rand.Intn(len(symbols))]
		)

		allAllowedCharacters += symbols
	}

	const remainingLenght = options.Lenght -  len(requiredCharacters)

	for i := 0; i < remainingLenght; i++{
		requiredCharacters = append(
			requiredCharacters,
			allAllowedCharacters[rand.Intn(len(allAllowedCharacters))]
		)
	}

	pw := utils.ShuffleBytes(requiredCharacters)

	return pw
}

