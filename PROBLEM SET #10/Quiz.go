// do it in csv format also make file in .csv
package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type quiz struct {
	question string
	option_a string
	option_b string
	option_c string
	option_d string
	answer   string
	//correct_answer string
}

func checkFileExistorNot() {
	/*check file exists we can use the Stat() and the isNotExists() function that the os package provides
	Stat() function is used to return the file info structure describing the file.
	*/
	if _, err := os.Stat("Q&A.txt"); err == nil {
		fmt.Printf("File exists\n")
	} else {
		fmt.Printf("File does not exist\n")
	}
}

/*
	func readFile() {
		fileName := "Q&A.txt"
		// read the whole content of file and pass it to file variable,
		//in case of error pass it to err variable
		file, err := ioutil.ReadFile("Q&A.txt")
		if err != nil {
			log.Panicf("failed reading data from file: %s", err)
		}
		fmt.Printf("\nFile Name: %s", fileName)
		fmt.Printf("\nSize: %d bytes", len(file))

//	fmt.Printf("\nData: %s \n", file)
}

	func quizQues&Ans()  {


		for i,j:=range quiz
		if quiz[i]==correctoption{
			fmt.Println("answers is correct")
			points++
		}
		else{
			fmt.Println("Wrong")
			fmt.Println("Correct option is ")
			points--
		}
	}
*/
func main() {
	fmt.Println("First Lets check file exists or not")
	checkFileExistorNot()
	fmt.Println("---------------------------------------------------")

	// open file
	f, err := os.Open("Q&A.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// convert records to array of structs
	quizrecord := quiz(data)

	// print the array
	fmt.Printf("%+v\n", quizrecord)

	/*	 fmt.Println("Read contents of file")
	readFile()
	fmt.Println("---------------------------------------------------")
	*/
}
