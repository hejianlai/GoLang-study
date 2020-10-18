package main

import `fmt`

type Person02 struct {
	 Age int
	 Name string
}
func main()  {
	var p1 Person02
	p1.Age = 123
	p1.Name = "xiaoming"
	var p2 *Person02 = &p1
	fmt.Println((*p2).Age)
	fmt.Println(p1.Name)
}
