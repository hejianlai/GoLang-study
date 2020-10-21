package main

import `fmt`
// 结构体
type Student struct {
	name string
	gender string
	age int
	id int
	score float64
}
// 结构体中声明一个say方法，返回string类型
func (student *Student) say() string  {
	infoStr := fmt.Sprintf("student的信息： name=[%v],gender=[%v],age=[%v],id=[%v],score=[%v]",
		student.name,student.gender,student.age,student.id,student.score)
	return infoStr
}
func main()  {
	// 创建Student结构体实例
	var stu = Student{
		name:   "hejianlai",
		gender: "3年级",
		age:    13,
		id:     2,
		score:  99.4,
	}
	// 访问say方法
	fmt.Println(stu.say())
}
