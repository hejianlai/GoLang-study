package main

import `fmt`

type Circle struct {
	redius float64
}
func (c Circle) area() float64  {
	return 3.14 * c.redius * c.redius
}
// 为了提高效率，通常我们方法和结构体的指针类型绑定
func (c *Circle) area2() float64  {
	// 因为c是指针，因此我们标准的访问其字段的方式是(*c).radius
	return 3.14 * c.redius * c.redius
}
func main()  {
	var c Circle
	c.redius = 4.0
	res := c.area()
	fmt.Println(res)
	res2 := (&c).area2()
	// 等价下面
	res3 := c.area2()
	fmt.Println(res2,res3)
}
