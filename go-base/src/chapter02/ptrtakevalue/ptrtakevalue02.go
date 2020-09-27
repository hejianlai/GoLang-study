package main

import `fmt`

func main()  {
	// 准备一个字符串类型
	var house  = "Malibe point 10880,90265"
	// 对字符串取地址，ptr类型为*string
	ptr := &house
	// 打印ptr的类型
	fmt.Printf("ptr type: %T\n",ptr)
	fmt.Printf("address: %p\n",ptr)
	value := *ptr
	fmt.Printf("value type: %T\n",value)
	fmt.Printf("value: %s\n",value)

}
