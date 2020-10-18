package main

import `fmt`

type Point struct {
	x int
	y int
}
type Rect struct {
	leftup,rightDown Point
}

func main()  {
	r1 := Rect{Point{1,2},Point{2,3}}
	fmt.Println(r1.leftup,r1.rightDown)
	fmt.Println(&r1.leftup.x,&r1.rightDown.y)
	r2 := Rect2
}
