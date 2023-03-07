package main

import (
	"flag"
	"fmt"

	"actest/tool"
)

const defaultInput = `{"name":"Default"}`

var input = flag.String("input", defaultInput, "the original string of input")

func main() {
	flag.Parse()

	fmt.Println(tool.ReadStringField(*input, "name"))
}
