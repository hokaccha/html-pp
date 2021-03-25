package main

import (
	"fmt"
	"os"

	"github.com/Joker/hpp"
)

func main() {
	s := hpp.Print(os.Stdin)
	fmt.Print(string(s))
}
