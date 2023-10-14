package main

import (
	_ "embed"
	"fmt"
)

//go:embed pg132.txt
var text []byte

func countByte(arr []byte, bitty byte) int {
	var count int = 0
	for _, elem := range arr {
		if elem == 10 {
			count++
		}
	}
	return count
}

func countWord(arr []byte) int {
	var count int = 0
	var prevElem byte = 10
	var elem byte
	for _, elem = range arr {
		if ((elem == 32) || (elem == 10) || (elem == 9) || (elem == 13)) && (prevElem != 9 && prevElem != 10 && prevElem != 32 && prevElem != 13) {
			count++
		}
		prevElem = elem
	}
	if !((elem == 32) || (elem == 10) || (elem == 9)) {
		count++
	}
	return count
}

func countChar(arr []byte) int {
	var count int = 0
	for i := 0; i < len(arr); i++ {
		if arr[i] < 192 {
			count++
		} else if arr[i] < 224 {
			count++
			i++
		} else if arr[i] < 240 {
			count++
			i += 2
		} else {
			count++
			i += 3
		}
	}
	return count
}

func main() {
	// fmt.Println(text)
	// Step One
	fmt.Println(len(text))
	// // Step Two
	lineCount := countByte(text, 10)
	fmt.Println(lineCount)
	//Step Three
	wordCount := countWord(text)
	fmt.Println(wordCount)
	// Step Four
	charCount := countChar(text)
	fmt.Println(charCount)
}
