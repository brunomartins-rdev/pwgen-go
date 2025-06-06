package utils

import (
	"errors"

	"github.com/brunomartins-rdev/pwgen-go/internal/generator"
)

func ValidateOptions(options generator.PasswordOptions) error {

	if (options.Lenght <= 0) {
		return errors.New("Password Lenght should be greater than 0")
	}

	enabledTypes := []bool{
		options.IncludeLowercase,
		options.IncludeUppercase,
		options.IncludeNumbers,
		options.IncludeSymbols,
	}

	count := 0
	for _, enabled := range enabledTypes {
		if enabled {
			count++
		}
	}

	if (count < options.Lenght) {
		return errors.New("You can't have more enabled parameters than lenght of password")
	}

	return nil
}
