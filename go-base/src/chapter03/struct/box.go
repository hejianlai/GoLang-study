package main

import `fmt`

type Box struct {
	len float64
	width float64
	height float64
}

func (box *Box) getVolumn() float64  {
	return  box.len * box.width * box.height
}
func main()  {
	var box =Box{
		len:    4.5,
		width:  5.3,
		height: 4,
	}
	fmt.Printf("%.2f",box.getVolumn())
}
