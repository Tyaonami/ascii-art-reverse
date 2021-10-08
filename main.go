package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//read font art
	fContent, err := os.ReadFile("standard.txt")
	check(err)
	fontData := string(fContent)
	font := strings.Split(fontData, "\n")
	//read art file for reverse
	textFile := os.Args[1][10:]
	textContent, err := os.ReadFile(textFile)
	check(err)
	textData := string(textContent)
	text := strings.Split(textData, "\n")

	if len(text) > 9 {
		for i := 0; i < len(text)-1; {

			reverse(font, text[i:i+8], 0, 0, 1)
			fmt.Println()
			i = i + 9

		}

	} else {
		reverse(font, text, 0, 0, 1)
		fmt.Println()
	}

}
func reverse(font []string, text []string, pos int, count int, start int) {
	if pos != len(text[count]) {
		l := len(font[start])
		if pos+l <= len(text[count]) {
			if count < 7 {
				if text[count][pos:l+pos] == font[start+count] {
					reverse(font, text, pos, count+1, start)
				} else {
					reverse(font, text, pos, 0, start+9)
				}
			} else {
				r := ((start - 1) / 9) + 32
				fmt.Printf("%c", r)
				reverse(font, text, pos+l, 0, 1)
			}
		} else {
			reverse(font, text, pos, 0, start+9)
		}
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
