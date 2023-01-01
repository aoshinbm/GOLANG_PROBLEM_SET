package main

import (
	"encoding/base64"
	"fmt"
)

func encode(str string) string {
	// base64.StdEncoding: Standard encoding with padding
	// It requires a byte slice so we cast the string to []byte
	encodedStr := base64.StdEncoding.EncodeToString([]byte(str))
	fmt.Println("Encoded:", encodedStr)
	return encodedStr
}

func decode(encodedStr string) {
	// Decoding may return an error, in case the input is not well formed
	decodedStr, err := base64.StdEncoding.DecodeString(encodedStr)
	if err != nil {
		panic("malformed input")
	}
	fmt.Println("Decoded:", string(decodedStr))

}
func main() {

	// String to encode
	str := "Encode exxample with Golang"
	encodedStr := encode(str)

	decode(encodedStr)

}
