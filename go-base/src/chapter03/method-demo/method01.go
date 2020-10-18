package main

import `fmt`

type Person struct {
	Name string
}
// func (p Person) test() {} 表示Person结构体有一个方法，方法名为test
// (p Person) 体现test方法是和A类型绑定的
// 给Perison类型绑定一个方法
func (p Person) test()  {
	fmt.Println("test()",p.Name)
}
func (p Person) speak()  {
	fmt.Println(p.Name,"是一个good man")
}
func (p Person) jisuan()  {
	res := 0
	for i := 1;i<1000;i++{
		res +=i
	}
	fmt.Println(p.Name,"计算的结果是：",res)
}
func (p Person) jisuan2(n int)  {
	res := 0
	for i := 1;i<n;i++{
		res +=i
	}
	fmt.Println(p.Name,"计算的结果是：",res)
}
// 需要返回值
func (p Person) getSum(n1 int,n2 int) int {
	return n1 + n2
}
func main()  {
	var p1 Person
	p1.Name = "tom"
	p1.test()
	p1.speak()
	p1.jisuan()
	p1.jisuan2(2)
	res := p1.getSum(10,20)
	fmt.Println("n1+n2的结果:",res)
}
// test方法和Person类型绑定
// test方法只能通过Person类型的变量来调用，而不能直接调用，也不能使用其他类型变量来调用
// func (p Person) test() {}...p 表示Person变量调用，这个p就是他的副本，这点和函数传参相似
// 这个名字，由程序员指定，不固定
