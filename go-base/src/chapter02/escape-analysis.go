package main

import `fmt`
// 本函数测试入口参数和返回值情况
func dummy(b int) int  {
	// 声明一个c赋值进入参数并返回
	var c int
	c = b
	return c
}
func void()  {

}
func main()  {
	var a int
	void()
	fmt.Println(a,dummy(3))
}
