package generator

import "math/rand"

type PasswordOptions struct {
	Lenght       int
	Uppercase    bool
	Lowercase    bool
	Numbers      bool
	Symbols      bool
	NumberOfGens int
}

const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!\"#$%&'()*+,-./:;<=>?@[\\]^_{|}~"

func GeneratePassword(options PasswordOptions) []string {
	list_pw := make([]string, options.NumberOfGens)

	for i := range list_pw {
		pw := make([]byte, options.Lenght) 

		// TODO: Ensure options outside of Lenght matter here
		for j := range pw {
			pw[j] = chars[rand.Intn(len(chars))]
		}

		list_pw[i] = string(pw)
	}
	
	return list_pw
}

