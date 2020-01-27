package main

import (
	"fmt"
)

func Factorial(n uint64)(result uint64)  {
	if(n>0){
		result = n* Factorial(n-1)
		return result
	}
	return 1
}
func main(){
	var i int = 20
	fmt.Printf("%d ,%d\n",i,Factorial(uint64(i)))
}