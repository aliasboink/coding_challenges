package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
)

func countByte(arr []byte, bitty byte) int {
	var count int = 0
	for _, elem := range arr {
		if elem == 10 {
			count++
		}
	}
	return count
}

// I can't bother to change this.
// It's a list of "separators" (tabs, whitesapces, newlines, etc.).
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
		switch {
		case arr[i] < 127:
			i += 0 //hehe
		case arr[i] >= 192 && arr[i] < 224:
			i += 1
		case arr[i] >= 224 && arr[i] < 240:
			i += 2
		case arr[i] >= 240 && arr[i] < 248:
			i += 3
		default:
			fmt.Println("Some funky characters snuck in!")
			os.Exit(1)
		}
		count++
	}
	return count
}

func readFile(fileName string) ([]byte, error) {
	text, err := os.ReadFile(fileName)
	if err != nil {
		return nil, errors.New("Issue with reading the specified file, fix it.")
	}
	return text, nil
}

func readStandardInput() []byte {
	scanner := bufio.NewScanner(os.Stdin)
	var text []byte
	for scanner.Scan() {
		text = append(text, scanner.Bytes()...)
		text = append(text, 10)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return text
}

func main() {
	bytesFlag := flag.Bool("c", false, "Count bytes")
	linesFlag := flag.Bool("l", false, "Count lines")
	wordsFlag := flag.Bool("w", false, "Count lines")
	charsFlag := flag.Bool("m", false, "Count lines")

	flag.Parse()
	args := flag.Args()

	var text []byte
	if len(args) == 0 {
		text = readStandardInput()
	} else if len(args) == 1 {
		var err error
		text, err = readFile(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		fmt.Println("Specify a file name or use the standard input!")
		os.Exit(1)
	}

	if *bytesFlag {
		fmt.Println(len(text))
	}
	if *linesFlag {
		lineCount := countByte(text, 10)
		fmt.Println(lineCount)
	}
	if *wordsFlag {
		wordCount := countWord(text)
		fmt.Println(wordCount)
	}
	if *charsFlag {
		charCount := countChar(text)
		fmt.Println(charCount)
	}
	if !(*charsFlag || *wordsFlag || *linesFlag || *bytesFlag) {
		lineCount := countByte(text, 10)
		wordCount := countWord(text)
		charCount := countChar(text)
		fmt.Println(lineCount, wordCount, charCount)
	}
}
