package main

import `fmt`

func main()  {
	angel := "Heros never die"
	angelBytes := []byte(angel)
	for i :=5; i <= 10 ; i++  {
		angelBytes[i] = ' '
	}
	fmt.Println(string(angelBytes))
}
