package main

import (
	`encoding/json`
	`fmt`
)

type A struct {
	Num int
}
type B struct {
	Num int
	Name string
}
type Monster struct {
	// `` 反射
	Name string `json:"name"`
	Age int `json:"age"`
	Skill string `json:"skill"`
}
func main()  {
	var a A
	var b B
	fmt.Println(a,b)
	monster := Monster{"牛魔王",15,"芭蕉扇"}
	fmt.Println(monster)
	// json序列化
	jsonStr,err := json.Marshal(monster)
	if err != nil{
		fmt.Println("报错的信息为:",err)
	}
	fmt.Println(string(jsonStr))
}

