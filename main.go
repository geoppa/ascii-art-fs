package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var input string
	var bannerName string

	// check if at least two arguments are provided
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run . <text> <banner>")
		return
	}

	// scan arguments to find the banner position (with or without .txt extension)
	bannerIndex := -1
	for i := 1; i < len(os.Args); i++ {
		arg := os.Args[i]
		// Aif the argument does not end with .txt, temporarily append it for validation
		if !strings.HasSuffix(arg, ".txt") {
			arg += ".txt"
		}

		if arg == "standard.txt" || arg == "shadow.txt" || arg == "thinkertoy.txt" {
			bannerName = arg // store the correct filename with .txt extension
			bannerIndex = i
			break
		}
	}

	// if no valid banner was found in the arguments, default to standard.txt and warn
	if bannerIndex == -1 {
		fmt.Println("Error: Invalid or missing banner. Only standard.txt, shadow.txt, and thinkertoy.txt are allowed. Using standard.txt instead.")
		bannerName = "standard.txt"

		// check if the user attempted a multi-word input even with a broken banner name
		if len(os.Args) > 3 {
			fmt.Println("Ungrouped, multi word text found. Printing only first word")
		}

		// since no valid banner matched, the first argument is assumed to be the text
		input = os.Args[1]
	} else {
		// check if there are ungrouped words between the first word and the banner
		if bannerIndex > 2 {
			fmt.Println("Ungrouped, multi word text found. Printing only first word")
		}
		// if a banner is found, print only the very first word (the first argument) with it
		input = os.Args[1]
	}

	// if the argument is "" exit without printing anything
	if input == "" {
		return
	}

	// read the file with the ascii art
	lines, err := readbanner(bannerName)
	if err != nil {
		fmt.Println("error reading file:", err)
		return
	}

	// split input in parts by newline characters
	parts := strings.Split(input, "\\n")

	// if the input is only "\n" print newlines only
	onlyNewlines := true
	for _, part := range parts {
		if part != "" {
			onlyNewlines = false
			break
		}
	}
	if onlyNewlines {
		for i := 0; i < len(parts)-1; i++ {
			fmt.Println()
		}
		return
	}

	for _, part := range parts {
		if part == "" {
			fmt.Println()
			continue
		}

		// print 8 vertical layers for the current part
		for i := 0; i < 8; i++ {
			for _, char := range part {
				// find the starting line index in the file for this character
				// ascii space is 32. each char block is 9 lines (8 art + 1 empty).
				startLine := int(char-32)*9 + 1

				// print the current line (i) of the character's art
				if startLine+i < len(lines) {
					fmt.Print(lines[startLine+i])
				}
			}
			fmt.Println() // move to the next vertical line of the ascii art
		}
	}
}

func readbanner(bannername string) ([]string, error) {
	file, err := os.Open(bannername)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// create an empty slice of strings (lines) to store the content of the file
	var lines []string

	// read through the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
