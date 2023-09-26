package main

import "fmt"

func main()  {
	//somethingelse("Hello World")
	//method 1
	//var mystate string
	//mystate = "Hello World"

	//method 2
    mystate := "Hello World"
	somethingelse(mystate)
}

func somethingelse(whattosay string)  {
	fmt.Println(whattosay)
}