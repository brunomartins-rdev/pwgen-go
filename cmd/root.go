package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/brunomartins-rdev/pwgen-go/config"
	"github.com/brunomartins-rdev/pwgen-go/internal/generator"
	"github.com/brunomartins-rdev/pwgen-go/internal/utils"
)

func Run() {

	lenghtFlag := flag.Int("lenght", config.DefaultLenght, "lenght of password")
	includeLowercaseFlag := flag.Bool("lowercase", config.DefaultIncludeLowercase, "include lowercase characters")
	includeUppercaseFlag := flag.Bool("uppercase", config.DefaultIncludeUppercase, "include lowercase characters")
	includeSymbolsFlag := flag.Bool("symbols", config.DefaultIncludeSymbols, "include symbols")
	includeNumbersFlag := flag.Bool("numbers", config.DefaultIncludeNumbers, "include numbers")

	flag.Parse()

	PwOptions := generator.PasswordOptions{
			Lenght: *lenghtFlag,
			IncludeLowercase: *includeLowercaseFlag,
			IncludeUppercase: *includeUppercaseFlag,
			IncludeNumbers: *includeNumbersFlag,
			IncludeSymbols: *includeSymbolsFlag,
	}

	err := utils.ValidateOptions(PwOptions)
	if (err != nil) {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	pw := generator.GeneratePassword(PwOptions)

	fmt.Println(pw)
}

