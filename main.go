package main

import (
	"math/rand"
	"time"

	"github.com/brunomartins-rdev/pwgen-go/cmd"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	cmd.Run()
}
