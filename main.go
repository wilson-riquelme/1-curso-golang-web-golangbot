package main

import "fmt"

func main() {  
	var width, height int = 100, 50 //declaring multiple variables
    var age int = 25 // variable declaration
    fmt.Println("My age is", age)
    age = 29 //assignment
    fmt.Println("My age is", age)
    age = 54 //assignment
    fmt.Println("My new age is", age , "height  is", height, "-" , width)
    //Variable tipo objeto
	var (
        name   = "naveen"
        age2    = 29
        height2 int
    )
    fmt.Println("my name is", name, ", age is", age2, "and height is", height2)
	//Short value 
	name2, age := "wilson", 29 //short hand declaration

    fmt.Println("my name is", name2, "age is", age)
}  