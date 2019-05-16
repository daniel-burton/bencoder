package main

import (
	"errors"
	"fmt"
	"strconv"
	//	"strings"
)

func main() {
	// test the Int decoder
	test := []string{"hi", "64", "ihelloe", "i32e", "i64e"}
	for _, v := range test {
		decode, err := DecodeInt(v)
		if err != nil {
			fmt.Printf("error: %v\n", err)
			continue
		}
		fmt.Printf("%T: %v\n", decode, decode)
	}
}

func DecodeInt(s string) (i int, err error) {
	// integers: i<integer in base ten ascii>e // note can include "-"
	if s[0] != 'i' || s[len(s)-1] != 'e' {
		err := errors.New(fmt.Sprintf("not a bencoded integer (must start with 'i' and end with 'e'): %q", s))
		return 0, err
	}
	i, err = strconv.Atoi(s[1 : len(s)-1])
	if err != nil {
		err = errors.New(fmt.Sprintf("error decoding %q as int", s[1:len(s)-1]))
	}
	return
}
