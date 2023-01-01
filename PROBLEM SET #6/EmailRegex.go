package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println(isEmailValid("test44@gmail.com"))    // true
	fmt.Println(isEmailValid("test$@gmail.com"))     // true -- expected "false"
	fmt.Println(isEmailValid("test44@gmail.com"))    // true
	fmt.Println(isEmailValid("bad-email"))           // false
	fmt.Println(isEmailValid("test44$@gmail.com"))   // false
	fmt.Println(isEmailValid("test-email.com"))      // false
	fmt.Println(isEmailValid("test+email@test.com")) // true
}

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[A-Za-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}
