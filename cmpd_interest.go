package main
import (
	"fmt"
	"math"
)

func main(){

	p:=1000.0
	r:=7.0
	t:=5.0
	c_i := p* (math.Pow((1 + r/100), t))
	fmt.Println(c_i)
}