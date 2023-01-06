package main
import(
	"fmt"
	"math"
)
func adv_CI(p,r,t,CI float64){
	
	CI = p* (math.Pow((1 + r/100), t))
	fmt.Println("Compound interest is: " ,CI)
}
func main(){
	var p,r,t float64
	fmt.Println("Enter the Principle amount: ")
	fmt.Scanln(&p)
	fmt.Println("Enter rate of interest: ")
	fmt.Scanln(&r)
	fmt.Println("Enter the time period: ")
	fmt.Scanln(&t)
	var CI float64
	CI=0
	adv_CI(p,r,t,CI)
}
