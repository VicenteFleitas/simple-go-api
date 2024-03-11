package main

import (
	"fmt"
	"strings"
)

func main() {
	var myString = []rune("resume")
	var indexed = myString[1]
	fmt.Printf("%v, %T\n", indexed, indexed)
	for i, v := range myString {
		fmt.Println(i, v)
	}
	fmt.Printf("The length of 'myString' is %v\n", len(myString))

	var myRune = 'a'
	fmt.Printf("myRune = %v\n", myRune)

	var strSlice = []string{"s", "u", "b", "s", "c", "r", "i", "b", "e"}
	var catStr = ""
	var strBuilder strings.Builder
	for i := range strSlice {
		// catStr += strSlice[i]
		strBuilder.WriteString(strSlice[i])
	}
	catStr = strBuilder.String()
	fmt.Printf("%v\n", catStr)
}
