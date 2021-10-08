package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//read font art
	fontFile := "shadow"
	fContent, err := os.ReadFile(fontFile + ".txt")
	check(err)
	fontData := string(fContent)
	font := strings.Split(fontData, "\n")
	//read file for reverse
	textFile := "test00"
	textContent, err := os.ReadFile(textFile + ".txt")
	check(err)
	textData := string(textContent)
	text := strings.Split(textData, "\n")
	reverse(font, text, 0, 0, 1)
	fmt.Println()

}
func reverse(font []string, text []string, pos int, count int, start int) {
	if pos != len(text[count]) {
		l := len(font[start])
		if pos+l <= len(text[count]) {
			//fmt.Println("pos", pos, pos+l)
			if count < 7 {
				//fmt.Println("count", count)
				if text[count][pos:l+pos] == font[start+count] {
					//fmt.Println("if")
					reverse(font, text, pos, count+1, start)
				} else {
					//fmt.Println("el")
					reverse(font, text, pos, 0, start+9)
				}
			} else {
				r := ((start - 1) / 9) + 32
				fmt.Printf("%c", r)
				//fmt.Println("newpos", pos+l)
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
