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

	tokens := clanglex.Lexicalize(string(input))

	fmt.Println(tokens)

}
