package main

import `fmt`

func main()  {
	// 设置元素数量为1000
	const elementCount =1000
	// 预分配足够多的元素切片
	srcData := make([]int,elementCount)
	// 将切片赋值
	for i:=0;i<elementCount ;i++  {
		srcData[i] = i
		// 引用切片数据
		refData := srcData
		// 预分配足够多的元素切片
		copyData := make([]int,elementCount)
		copy(copyData,srcData)
		srcData[0] = 999
		fmt.Println(refData[0])
		fmt.Println(copyData[0],copyData[elementCount-1])
		copy(copyData,srcData[4:6])
		for i:=0; i<5;i++  {
			fmt.Printf("%d",copyData[i])
		}
	}
}
