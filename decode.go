package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	test := "i100ei200e6:aaaaaaimistakei64e4:abcd13:abcdefghijklm"
	Parse(test)
}

func Parse(s string) { // return type will be []interface{}
	//
	var r *strings.Reader = strings.NewReader(s)
	var b strings.Builder
	var inString bool
	var inInt bool
	var inLen bool
	// var listLevel int
	// var dictLevel int
	var sLen int

	for {
		c, _, err := r.ReadRune()
		fmt.Printf("%q\n", c)
		if err != nil {
			fmt.Printf("Reader error: %v\n", err)
			break
		}

		if inString == false && inInt == false && inLen == false { //if totally outside
			if unicode.IsDigit(c) { //if first rune is a number, its a string length
				inLen = true
				fmt.Print("\tnow in len\n")
				if _, err := b.WriteRune(c); err != nil { //write rune to builder
					fmt.Printf("Writer error: %v", err)
				}
			} else if c == 'i' {
				inInt = true
				fmt.Print("\tnow in Int\n")
				continue
			}
		} else if inInt == true {
			if c == 'e' {
				inInt = false
				i, err := strconv.Atoi(b.String())
				if err != nil {
					e := fmt.Sprintf("error decoding %q as int", b.String())
					fmt.Printf("%v\terror: %v\n", e, err)
					b.Reset()
					continue
				}
				fmt.Printf("--------------------int: %v\n", i)
				b.Reset()
			} else {
				if _, err := b.WriteRune(c); err != nil { //write rune to builder
					fmt.Printf("Writer error: %v", err)
				}
			}
		} else if inLen == true {
			if c == ':' {
				inLen = false
				sLen, err = strconv.Atoi(b.String())
				if err != nil {
					e := fmt.Sprintf("error decoding %q as int", b.String())
					fmt.Printf("%v\terror: %v\n", e, err)
				}
				inString = true
				fmt.Println("\tNow in string")
				b.Reset()
				continue
			} else {
				if _, err := b.WriteRune(c); err != nil {
					fmt.Printf("Writer error: %v", err)
				}
			}
		} else if inString == true {
			if _, err := b.WriteRune(c); err != nil {
				fmt.Printf("Writer error: %v", err)
			}
			sLen -= 1
			if sLen == 0 {
				fmt.Printf("--------------------string: %v\n", b.String())
				inString = false
				b.Reset()
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
	i, err = strconv.Atoi(s)
	if err != nil {
		err = errors.New(fmt.Sprintf("error decoding %q as int", s))
	}
	return
}
