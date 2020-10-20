package main

import (
	`fmt`
)

type MethodUtis struct {

}
type Calcuator struct {
	Num1 float64
	Num2 float64
}

func (calcuator *Calcuator) getSum() float64  {
	return calcuator.Num1 + calcuator.Num2
}
func (calcuator *Calcuator) getRes(operator byte) float64 {
	res := 0.0
	switch operator {
	case '+':
		res = calcuator.Num1 + calcuator.Num2
	case '-':
		res = calcuator.Num1 - calcuator.Num2
	case '*':
		res = calcuator.Num1 * calcuator.Num2
	case '/':
		res = calcuator.Num1 / calcuator.Num2
	default:
		fmt.Println("运算符输入有误。。。")
	}
	return res
}
func (mu MethodUtis) Print() {
	for i :=1;i<=10;i++{
		for j :=1;j<=8;j++{
			fmt.Print("*")
		}
		fmt.Println()
	}
}
func (mu MethodUtis) Print2(m int,n int) {
	for i :=1;i<=m;i++{
		for j :=1;j<=n;j++{
			fmt.Print("*")
		}
		fmt.Println()
	}
}
func (mu *MethodUtis) Arae(len float64,width float64) (float64) {
	return len * width
}
func (mu *MethodUtis) JudgeNum(num int)  {
	if num % 2 == 0{
		fmt.Println(num,"是偶数")
	}else {
		fmt.Println(num,"是奇数")
	}
}
func (mu *MethodUtis) Print3(n int,m int,key string)  {
	for i :=1;i<=n;i++{
		for j:=1;j<=m;j++{
			fmt.Print(key)
		}
		fmt.Println()
	}
}
func main()  {
	var mu MethodUtis
	mu.Print()
	fmt.Println()
	mu.Print2(3,5)
	arae := mu.Arae(34.2,4.4)
	fmt.Println(arae)
	mu.JudgeNum(21)
	mu.Print3(3,6,"$")
	var calcuator Calcuator
	calcuator.Num1 = 1
	calcuator.Num2 = 3
	fmt.Println(calcuator.getSum())
	res := calcuator.getRes('-')
	fmt.Println("res=",res)
}
