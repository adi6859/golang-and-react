package main
import(
	"fmt"
	"math"
	
	
)
func main(){
	var n int
	fmt.Println("Enter number: ")
	fmt.Scanln(&n)
	c:=0
	for n>0{
		n=n/10
		c++
	}
	fmt.Println(c)
	fmt.Println(math.Floor(math.Log10(float64(n))+1))


}