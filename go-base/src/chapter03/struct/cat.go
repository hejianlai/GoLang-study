package main

import `fmt`

type Cat struct {
	Name string
	Age int
	Sex string
	Hobby string
}

func main()  {
	var cat1 Cat
	var cat2 Cat
	cat1.Name = "笑话"
	cat1.Age = 123
	cat1.Sex = "男"
	cat1.Hobby = "吃鱼"
	fmt.Println("名字：",cat1.Name)
	fmt.Println("年龄：",cat1.Age)
	fmt.Println("性别：",cat1.Sex)
	fmt.Println("好爱：",cat1.Hobby)
}
