package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println(passwordValid("abc"))      // true
	fmt.Println(passwordValid("abc123"))   // true -- expected "false"
	fmt.Println(passwordValid("abC12312")) // true
	fmt.Println(passwordValid("Aoshin"))   // false
	fmt.Println(passwordValid("28Aoshin")) // false
	fmt.Println(passwordValid("7#Aoshin")) // false
}

func passwordValid(e string) bool {
	passwordRegex := regexp.MustCompile(`^(?=.*[0-9])(?=.*[a-z])(?=.*[A-Z])(?=.*[*.!@$%^&(){}[]:;<>,.?/~_+-=|\]).{8,32}$`)

	/*
		Dissecting the pattern
		^                                            Match the beginning of the string
		(?=.*[0-9])                                  Require that at least one digit appear anywhere in the string
		(?=.*[a-z])                                  Require that at least one lowercase letter appear anywhere in the string
		(?=.*[A-Z])                                  Require that at least one uppercase letter appear anywhere in the string
		(?=.*[*.!@$%^&(){}[]:;<>,.?/~_+-=|\])    Require that at least one special character appear anywhere in the string
		.{8,32}                                      The password must be at least 8 characters long, but no more than 32
		$                                            Match the end of the string.
	*/
	return passwordRegex.MatchString(e)
}
