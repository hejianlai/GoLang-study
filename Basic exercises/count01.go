package main
import "fmt"
func main(){
	var a int = 21
	var b int = 33
	var c int
	c = a % b
	fmt.Printf("%d", c )
	if (a<b){
		fmt.Printf("true")
	} else {
		fmt.Printf("false")
	}
	var e bool = true
	var f bool = false
	if(e&&f){
		fmt.Printf("1")
	}
	if (e||f){
		fmt.Printf("2")
	}
}