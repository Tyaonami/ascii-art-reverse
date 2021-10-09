package main

import (
	"fmt"
	"os"
	"strings"
)

func tttt() {
	//read font art
	fontFile := "shadow"
	fContent, err := os.ReadFile(fontFile + ".txt")
	check1(err)
	fontData := string(fContent)
	font := strings.Split(fontData, "\n")
	//read file for reverse
	testFile := "test00"
	testContent, err := os.ReadFile(testFile + ".txt")
	check1(err)
	testData := string(testContent)
	test := strings.Split(testData, "\n")
	fmt.Println(len(test[0]), test[0])
	pos := 0
	count := 0
	var word []int
	for i := 1; i < len(font); i = i + 9 {
		l := len(font[i])
	out:
		for j := 0; j < 9; j++ {
			if len(test[j]) >= l+pos {
				if test[j][pos:l+pos] == font[i+j] {
					count = count + 1
				} else {
					count = 0
					break out
				}
			}
		}
		//fmt.Println("Count ", count, i)
		if count == 8 {
			word = append(word, i)
			r := ((i - 1) / 9) + 32
			fmt.Printf("%c", r)
			count = 0
			pos = pos + l
			if pos != len(test[0]) {
				i = -8
			}
		}
	}
	fmt.Println()
}
func check1(e error) {
	if e != nil {
		panic(e)
	}
}
