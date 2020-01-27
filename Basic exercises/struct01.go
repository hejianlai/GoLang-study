package main

import "fmt"

type Books struct {
	title string
	author string
	subject string
	book_id string
}
func main(){
	fmt.Println(Books{"GO langage","www.runoob.com","go stu","321"})
}