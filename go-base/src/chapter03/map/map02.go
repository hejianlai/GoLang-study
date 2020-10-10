package main

import `fmt`

func main()  {
	var numbers []int
	for i:=0;i<30;i++{
		numbers = append(numbers,i)
		fmt.Printf("len: %d cap: %d pointer: %p\n",len(numbers),cap(numbers),numbers)
	}
}
