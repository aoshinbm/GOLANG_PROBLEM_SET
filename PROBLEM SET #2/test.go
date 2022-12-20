package main
import "fmt"
 
func main(){	
	calling()
}
 
func calling(){
 
	for(checkFlagExit){
		var inputFloatValue string
		fmt.Println("For Exit press \"q\" or \"Q\")
		fmt.Scanln(&inputFloatValue)
		sliceOfFloatingPoint(inputFloatValue)
	}
 
}
 
func sliceOfFloatingPoint(input){
	slicePoint := sliceofPoints //v need to create a blank space
	if(input == q || input == Q){
		checkFlagExit = false
		sd(sliceOfFloat)
	} else {
		slicePoint.append(input)
	}
}
 
 
 
 
 
func sd(x string){
 
	var floatNum float32
	var sum float32
 
 
	sum = 0
	checkValue := 'q'
	str := x
 
 
	for checkValue != x{
		fmt.Println("Enter Floating number :")
		fmt.Scan(&floatNum)
		sum = sum + floatNum
		calling()
	}
	fmt.Println(sum)
}