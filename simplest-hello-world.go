/* this is the simplest Hello World program in Go 
It does not use the "fmt" library which is confusing for the newbies 
and completely unnecessary for printing "Hello World" */


package main  //mandatory. indicates that this is the main module. This is the file where the execution of the program starts

func main() {     //defining the "main" function. This is also mandatory in every program that is executed. It marks where the program's execution starts
	println("Hello World!")      //the "println" built-in function prints out to the console the string (or whatever parameter it is given) and adds a new line.
}
