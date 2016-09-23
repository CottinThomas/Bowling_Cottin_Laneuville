package main

import "fmt"

func GetString() string {
	//return "azerty" // test fail
	return "This is a string" // test pass
}

func main() {
	fmt.Println("I got the string \""+GetString()+"\"!")	
}