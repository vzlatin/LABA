package main

import (
	"fmt"
	"os"

	"github.com/vzlatin/LABA/repl"
)

func main() {
	fmt.Println("Hai salut frumosule! Te invit la o LABA")
	repl.Start(os.Stdin, os.Stdout)
}
