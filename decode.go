package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	test := "i100ei200e6:ihulloei64e4:13:"
	Parse(test)
}

func Parse(s string) { //what should the return type be?
	//
	var r *strings.Reader = strings.NewReader(s)
	var b strings.Builder
	var inString bool
	var inInt bool
	var inLen bool
	// var sLen int

	for {
		c, _, err := r.ReadRune()
		fmt.Printf("\t%q\n", c)
		if err != nil {
			//fmt.Printf("Reader error: %v\n", err)
			break
		}

		if inString == false && inInt == false && inLen == false { //if totally outside
			if unicode.IsDigit(c) { //if first rune is a number, its a string length
				inLen = true
				fmt.Print("now in len\n")
				if _, err := b.WriteRune(c); err != nil { //write rune to builder
					fmt.Printf("Writer error: %v", err)
				}
			} else if c == 'i' {
				inInt = true
				fmt.Print("now in Int\n")
				continue
			}
		} else if inInt == true {
			if c == 'e' {
				inInt = false
				i, err := strconv.Atoi(b.String())
				if err != nil {
					fmt.Sprintf("error decoding %q as int", s)
					fmt.Printf("error: %v\n", err)
					b.Reset()
					continue
				}
				fmt.Printf("int: %v\n", i)
				b.Reset()
			} else {
				if _, err := b.WriteRune(c); err != nil { //write rune to builder
					fmt.Printf("Writer error: %v", err)
				}
			}
		} else if inLen == true {
			if c == ':' {
				inLen = false
				// sLen, err = strconv.Atoi(b.String())
				// if err != nil {
				// 	fmt.Sprintf("error decoding %q as int", s)
				// 	fmt.Printf("error: %v\n", err)
				// }
				fmt.Printf("string length, %q\n", b.String())
				// inString = true
				b.Reset()
				continue
			} else {
				if _, err := b.WriteRune(c); err != nil {
					fmt.Printf("Writer error: %v", err)
				}
			}
		}

	}
}

func DecodeList(s string) /*(o string, err error)*/ {
	fmt.Printf("list to decode: %q", s)
}

func DecodeDict(s string) /*(o string, err error)*/ {
	fmt.Printf("dict to decode: %q", s)
}

func DecodeInt(s string) (i int, err error) {
	// integers: i<integer in base ten ascii>e // note can include "-"
	// if s[0] != 'i' || s[len(s)-1] != 'e' {
	// 	err := errors.New(fmt.Sprintf("not a bencoded integer (must start with 'i' and end with 'e'): %q", s))
	// 	return 0, err
	// }
	i, err = strconv.Atoi(s)
	if err != nil {
		err = errors.New(fmt.Sprintf("error decoding %q as int", s))
	}
	return
}
