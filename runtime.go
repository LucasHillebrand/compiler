package main

import (
	"fmt"
	"os"
	"os/exec"
)

type pack struct {
	Name       string
	CompileCmd string
	RunCmd     string
}

var globalPack pack

func runtimeEnv(command string, args []string, mode int) {
	switch command {
	case "name":
		globalPack.Name = args[0]
	case "compiler":
		setCmds(args[0])
	case "build":
		build(args[0])
	case "copyFile":
		copyFile(args, mode)
	case "createDir":
		createDir(args[0])
	}
}

func createDir(name string) {
	os.Mkdir(name, os.FileMode(0777))
}

func copyFile(args []string, mode int) {
	if mode >= 2 {
		data, _ := os.ReadFile(args[0])
		os.WriteFile(args[1], data, os.FileMode(0666))
	}
}

func setCmds(arg string) {
	csv := getCsvFile("./compiler.csv")
	line := searchLine(arg, 0, csv)
	if line == -1 {
		fmt.Println("! >> error compiler not found << !")
		os.Exit(255)
	} else {
		globalPack.CompileCmd = csv[line][1]
		globalPack.RunCmd = csv[line][2]
	}
}

func build(arg string) {
	command := globalPack.CompileCmd
	tmpcmd := split(command, "%s")
	if len(tmpcmd) == 2 {
		command = tmpcmd[0] + arg + tmpcmd[1]
	}

	tmpcmd = split(command, "%o")
	if len(tmpcmd) == 2 {
		command = tmpcmd[0] + globalPack.Name + tmpcmd[1]
	}

	cmd := exec.Command("bash", "-c", command)
	cmd.Run()
}

/*func print2Dlist(test [][]string) {
	for i := 0; i < len(test); i++ {
		for j := 0; j < len(test[i]); j++ {
			fmt.Printf("|> %s <|", test[i][j])
		}
		fmt.Printf("\n")
	}
}*/
