package main

import (
	"fmt"
	"os"
)

func main() {
	arg := 0
	if len(os.Args) > 1 {
		fmt.Println(len(os.Args))
		arg = getMode(os.Args[1])
	}
	file := readLines("./.main.cfg")
	for i := 0; i < len(file); i++ {
		tmpLst := split(file[i], ":")
		if len(tmpLst) > 1 {
			runtimeEnv(tmpLst[0], getArgs(tmpLst[1]), arg)
		}
	}
}

func getMode(inp string) int {
	var out int
	switch inp {
	case "build":
		out = 1
	case "install":
		out = 2
	}
	return out
}
