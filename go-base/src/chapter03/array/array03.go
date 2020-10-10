package main

import (
	`fmt`
)

func main()  {
	var  strList []string
	var numList []int
	var numListEmpty = []int{}
	fmt.Println(strList,numList,numListEmpty)
	fmt.Println(len(strList),len(numListEmpty),len(numList))

	fmt.Println(strList == nil)
	fmt.Println(numList == nil)
	fmt.Println(numListEmpty == nil)
}
