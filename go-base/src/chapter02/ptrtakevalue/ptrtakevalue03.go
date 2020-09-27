package main

import `fmt`

func swap(a,b *int)  {
	// 取a指针的值，赋给临时变量t
	t := *a
	// 取b指针的值，赋给a指正指向的变量
	*a = *b
	// 将a指针的值赋给b指针指向的变量
	*b = t
}
func main()  {
	// 准备两个变量,赋值1和2
	x ,y := 1, 332
	// 交换变量值
	swap(&x,&y)
	// 输出变量值
	fmt.Println("a,b=", x,y)
}