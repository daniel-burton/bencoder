package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	// test the Int decoder
	//test := []string{"hi", "64", "ihelloe", "i32e", "i64e"}
	test := "i100ei200eihulloei64e"
	Parse(test)
	// for _, v := range test {
	// 	decode, err := DecodeInt(v)
	// 	if err != nil {
	// 		fmt.Printf("error: %v\n", err)
	// 		continue
	// 	}
	// 	fmt.Printf("%T: %v\n", decode, decode)
	//}
}

func Parse(s string) { //what should the return type be?
	//
	var r *strings.Reader = strings.NewReader(s)
	var b strings.Builder
	var inString bool
	var inInt bool
	//var inLen bool
	//var sLen int

	for {
		c, _, err := r.ReadRune()
		if err != nil {
			fmt.Printf("Reader error: %v", err)
			break
		}

		if inString == false && inInt == false { //if totally outside
			if unicode.IsDigit(c) { //if first rune is a number, its a string length
				inString = true
				//sLen = int(c)
			} else if c == 'i' {
				inInt = true
				continue
			}
		} else if inInt == true {
			if c == 'e' {
				inInt = false
				i, err := DecodeInt(b.String())
				if err != nil {
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
		} //else if inString == true {

		//}
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
