package main

import `fmt`

func main()  {
	var a map[string]string
	a = make(map[string]string,10)
	a["no1"] = "宋江"
	a["no2"] = "吴勇"
	fmt.Println(a)

	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "上海"
	fmt.Println(cities)
}
