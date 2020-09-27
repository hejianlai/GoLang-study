package main

import (
	`fmt`
)

func main()  {
	theme := "阻击战 start"
	fmt.Println(len(theme))
	for _, s := range theme{
		fmt.Printf("Unicode: %c %d\n",s,s)
	}
}
