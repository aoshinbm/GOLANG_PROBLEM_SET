package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	//check first file file exist or not
	//Stat() function is used to return the file info structure describing the file
	if _, err := os.Stat("LearnerData.txt"); err == nil {
		fmt.Println("File exists...")
	} else {
		fmt.Println("File does not exist!!!!!")
	}

	fmt.Println("----------------------------------------------")
	//read data from the file
	fmt.Println("Reading from file")
	fileName := "LearnerData.txt"

	// ioutil package contains inbuilt methods like ReadFile that reads the
	// filename and returns the contents.
	data, err := ioutil.ReadFile("LearnerData.txt")
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}
	fmt.Printf("\nFile Name: %s", fileName)
	fmt.Printf("\nSize: %d bytes", len(data))
	fmt.Printf("\nData:\n %s", data)

	fmt.Println("\n")
	fmt.Println("----------------------------------------------")
	fmt.Println("\n")

	//Use a bufio.Scanner to read a file line by line.
	file, err := os.Open("LearnerData.txt")
	//os.open() function to open the file.

	if err != nil {
		log.Fatal(err)
	}
	//Defer is used to ensure that a function call is performed later in a program’s execution,
	//usually for purposes of cleanup
	defer file.Close()

	//bufio.NewScanner() function to create the file scanner
	scanner := bufio.NewScanner(file)
	//scanner Scan() function is used in a for loop to get each line and process it
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
