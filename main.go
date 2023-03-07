package main

import (
	"flag"
	"fmt"

	"actest/tool"
)

const defaultInput = `{"name":"Default"}`

var input = flag.String("input", defaultInput, "the original string")

func main() {
	flag.Parse()

	fmt.Println(tool.ReadStringField(*input, "name"))
}
