package main

import `fmt`

func main()  {
	seq := []string{"a","b","c","d","e"}
	// 指定删除位置
	index := 2
	// 查看删除位置之前的元素和之后的元素
	fmt.Println(seq[:index],seq[index+1: ])
	// 将删除点前的元素连接起来
	seq = append(seq[:index],seq[index+1:]...)
	fmt.Println(seq)
}