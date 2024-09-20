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
			s = filepath.Base(os.Args[0])
		} else {
			s += sep + os.Args[i]
		}
	}
	fmt.Println(s)
}
