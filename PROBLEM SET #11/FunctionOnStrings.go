package main

import (
	"fmt"
	"strings"
)

func characterFilter(str []string, myFun func(string) bool) string {
	sliceOfString := make([]string, 0)
	for _, strrings := range str {
		if myFunc(strrings) {
			sliceOfString = append(sliceOfString, strrings)
		}
	}
	return sliceOfString
}

func main() {

	str := []string{"ant", "beetle", "bee", "wasp", "butterfly", "fly", "moth"}
	fmt.Println("---------------------------------------")

	wordsWith3Letters := characterFilter(str, func(strr []string) bool {
		return (len(strrings) == 3) == 0
	})
	fmt.Println(wordsWith3Letters)

	fmt.Println("---------------------------------------")

	str2 := "b"
	wordsWithB := characterFilter(string1, func(strr []string) bool {
		return strings.Contains(string1, str2) == 0
	})
	fmt.Println(wordsWith3Letters)

}

/* function  -- ()

func functionName (paramtere list -- Any data type(User defiend or golang provide) ) retun type {

}

Higher order function --- () --- Abstracting the logic

func functionName (paramtere list -- another func or func's) ) retun func {

}



*/
