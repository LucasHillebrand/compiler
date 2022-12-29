package main

import (
	"os"
)

func count(data, sep string) int {
	counter := 1

	for i := 0; i < len(data); i++ {
		if nextChars(data, i, len(sep)) == sep {
			counter++
		}
	}
	return counter
}

func readLines(path string) []string {
	data, err := os.ReadFile(path)
	if err != nil {
		return []string{"error"}
	}
	datastr := ""
	for i := 0; i < len(data); i++ {
		datastr += string(data[i])
	}
	return split(datastr, "\n")
}

func split(data, seperator string) []string {
	out := make([]string, count(data, seperator))
	out[0] = ""
	for i, col := 0, 0; i < len(data); i++ {
		if nextChars(data, i, len(seperator)) == seperator {
			i += len(seperator) - 1
			col++
			out[col] = ""
		} else {
			out[col] += string(data[i])
		}
	}
	return out
}

func nextChars(data string, index, lenght int) string {
	out := ""
	for i := index; i < index+lenght && i < len(data); i++ {
		out += string(data[i])
	}
	return out
}

func getArgs(argString string) []string {
	length := count(argString, "\"") / 2
	out := make([]string, length)
	for index, item, isActive := 0, 0, false; index < len(argString); index++ {
		if isActive && string(argString[index]) == "\"" {
			isActive = false
			if item+1 < length {
				item++
			}
		} else if !isActive && string(argString[index]) == "\"" {
			isActive = true
		} else if isActive {
			out[item] += string(argString[index])
		}
	}
	return out
}

//csvfile

func getCsvFile(path string) [][]string {
	csvRaw := readLines(path)
	out := make([][]string, len(csvRaw))
	for i := 0; i < len(csvRaw); i++ {
		out[i] = split(csvRaw[i], ",")
	}
	return out
}

func searchLine(keyword string, column int, data [][]string) int {
	out := -1
	for i := 0; i < len(data); i++ {
		if data[i][column] == keyword {
			out = i
			i = len(data)
		}
	}
	return out
}
