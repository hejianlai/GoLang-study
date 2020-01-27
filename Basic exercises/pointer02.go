package main

import (
	"fmt"
)

func main()  {
	var a int = 10
	var ip *int
	ip =&a
	fmt.Printf("%x\n",&a)
	fmt.Printf("ip:%x\n",ip)
	fmt.Printf("*ip:%d\n",*ip)
}
