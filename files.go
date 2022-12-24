package main

import "os"

func printStrLst(data []string) string {
	out := ""
	for i := 0; i < len(data); i++ {
		if i == len(data)-1 {
			out += data[i]
		} else {
			out += data[i] + " ||"
		}
	}
	return out
}
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
