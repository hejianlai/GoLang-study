package main

import (
	`fmt`
)

func main()  {
	var monsters []map[string]string
	monsters = make([]map[string]string,2)
	if monsters[0] == nil{
		monsters[0] = make(map[string]string,2)
		monsters[0]["name"] = "牛魔王"
		monsters[0]["age"] = "5000"

	}
	if monsters[1] == nil{
		monsters[1] = make(map[string]string,2)
		monsters[1]["name"] = "牛魔王"
		monsters[1]["age"] = "5000"
		monsters[1]["addr"] = "火焰山"


	}
	newmonster := map[string]string{
		"name" : "火云邪神",
		"age" : "234",

	}
	monsters = append(monsters,newmonster)
	fmt.Println(monsters)
}
