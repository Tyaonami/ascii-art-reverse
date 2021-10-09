package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	ok, er := checkArgs() // the function checks the validity of the arguments
	if ok == true {       // if we get from checkArgs true, work starts =)
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
				if len(text[i]) > 0 {
					reverse(font, text[i:i+8], 0, 0, 1)
					fmt.Println()
					i = i + 8
				} else {
					fmt.Println()
					i = i + 1
				}
			}
		} else {
			reverse(font, text, 0, 0, 1)
			fmt.Println()
		}
	} else {
		fmt.Print(er)
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

// the last argument must be started from --output= and contains at least 1 character
func checkArgs() (bool, string) {
	ok := true
	s := ""
	if len(os.Args) != 2 || len(os.Args[1]) < 9 || os.Args[1][0:10] != "--reverse=" {
		ok = false
		s = "Usage: go run . [OPTION]\n\nEX: go run main.go --reverse=<fileName>\n"
	}
	return ok, s
}
