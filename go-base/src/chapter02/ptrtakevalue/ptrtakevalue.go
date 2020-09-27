package main

import (
	`fmt`
)

func main()  {
	var cat int = 123
	var str string = "banana"
	fmt.Printf("%p %p",&cat,&str)
}
