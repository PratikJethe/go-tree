package main

import (
	"fmt"

	"github.com/pratikjethe/go-tree/cmd"
	"github.com/pratikjethe/go-tree/tree"
)

func main() {
	
	inputFlags := cmd.GetInput()
	output := tree.GetOutput(inputFlags)
	fmt.Println(output)
}
