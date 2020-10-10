package main

import `fmt`

func main()  {
	var highRiseBuilding [39]int
	for i := 0; i< 39;i++{
		highRiseBuilding[i] = i + 1
	}
	// 区间
	fmt.Println(highRiseBuilding[10:15])
	// 中间到尾部的所有元素
	fmt.Println(highRiseBuilding[20:])
	// 开头到中间的所有元素
	fmt.Println(highRiseBuilding[:6])
}

