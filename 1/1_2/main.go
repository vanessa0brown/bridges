package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	var s, sep string
	for i := range len(os.Args) {
		sep = " "
		if i == 0 {
			fmt.Println("Название программы:", filepath.Base(os.Args[0]))
		} else {
			fmt.Printf("[%v] %v \n", i, sep+os.Args[i])
		}
	}
	fmt.Println(s)
}
