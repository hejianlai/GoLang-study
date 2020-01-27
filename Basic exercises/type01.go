package main
import "fmt"
var a int = 20;
func main(){
	var a int =10
	var b int =20
	var c int =0
	fmt.Printf("main()函数中的a=%d\n",a);
	c = sum(a,b);
	fmt.Printf("main()函数中的c=%d\n",c);
	
}

func sum(a,b int) int{
	fmt.Printf("sum()函数中 a=%d\n",a);
	fmt.Printf("main()函数中c = %d\n",b);
	return a + b;
	
}