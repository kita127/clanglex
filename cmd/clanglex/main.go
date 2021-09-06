package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/kita127/clanglex"
)

func main() {

	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	tokens, err := clanglex.Lexicalize(string(input))
	if err != nil {
		panic(err)
	}

	fmt.Println(tokens)

}
