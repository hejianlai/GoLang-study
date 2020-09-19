package main

import (
	`fmt`
	`math`
)

func main()  {
	// 输出各数值范围
	fmt.Println("int8 range",math.MinInt8,math.MaxInt8)
	fmt.Println("int16 range",math.MinInt16,math.MaxInt16)
	fmt.Println("int32 range",math.MinInt32,math.MaxInt32)
	fmt.Println("int64 range",math.MinInt64,math.MaxInt64)
	// 初始化一个32位整型值
	var a int32 = 1023533242
	// 输出变量数值转换为十六进制，发生数值截断
	fmt.Printf("int32 0x%x %d\n",a,a)
	 b := int16(a)
	 fmt.Printf("int16: 0x%x %d\n",b,b)
	 var c float32	= math.Pi
	 fmt.Println(int(c))
}
