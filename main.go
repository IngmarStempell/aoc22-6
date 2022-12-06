package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fileHandle, _ := os.Open("input.txt")
	reader := bufio.NewReader(fileHandle)
	line, _ := reader.ReadString('\n')

	leng := 14
	for i := 0; i < len(line)-leng; i++ {
		fmt.Println(line[i : i+leng])
		if allCharsDiff(line[i : i+leng]) {
			fmt.Println(i + leng)
			break
		}
	}
}
func allCharsDiff(input string) bool {
	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			if input[i] == input[j] {
				return false
			}
		}
	}
	return true
}
