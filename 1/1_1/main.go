package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var s, sep string
	for i := range len(os.Args) {
		sep = " "
		if i == 0 {
			idx = strings.LastIndex(os.Args[0], "/")
			s := os.Args[0]
			s = strings.Slice(os.Args)
		} else {
			s += sep + os.Args[i]
		}
	}
	fmt.Println(s)
}
