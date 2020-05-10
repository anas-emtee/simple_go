package main

import (
	"fmt"

	"github.com/dgreat91/simple_go/octopus"
	"github.com/dgreat91/simple_go/stringutil"
)

func main() {
	fmt.Printf("hello, world\n")
	fmt.Println(stringutil.Reverse("!oG ,ereht olleH") + "\n")

	oct := octopus.Octopus{
		Name:  "Jesse",
		Color: "orange",
	}

	fmt.Println(oct.String())

	oct.reset()
}
