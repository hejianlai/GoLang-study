package main

import (
	`fmt`
	`sort`
)

func main()  {
	map1 := make(map[int]int)
	map1[0] = 123
	map1[23] = 23
	map1[2] = 233
	map1[22] = 2123
	fmt.Println(map1)

	var keys []int
	for k, _ := range map1  {
		keys=append(keys,k)
	}
	sort.Ints(keys)
	fmt.Println(keys)
	for _, k := range keys{
		fmt.Printf("map[%v]=%v\n",k,map1[k])
	}
}
