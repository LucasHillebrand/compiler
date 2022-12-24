package main

import (
	"fmt"
)

func main() {
	file := readLines("./.main.cfg")
	for i := 0; i < len(file); i++ {
		tmpLst := split(file[i], ":")
		fmt.Println(printStrLst(tmpLst))
	}
}
