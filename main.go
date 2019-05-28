package main

import "fmt"

func main() {
	str := displayHello()
	fmt.Println(str)
}

func displayHello() string {
	return "Hello Go"
}
