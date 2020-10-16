package main

import `fmt`

type Person struct {
	Name string
	Age int
	Scores [5]float64
	ptr *int
	slice []int
	map1 map[string]string
}
type Student struct {
	Name string
	Age int
}
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
