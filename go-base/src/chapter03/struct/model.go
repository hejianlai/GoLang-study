package main

import (
	`fmt`
)

func main()  {
	var p1 Person
	stu1 := Student{"xiaomiang",23}
	fmt.Println(stu1)
	fmt.Println(p1)
	if p1.ptr == nil{
		fmt.Println(p1.ptr)
		fmt.Println("ok1")
	}
	if p1.Age == 4{
		fmt.Println("ok2")
	}

}
