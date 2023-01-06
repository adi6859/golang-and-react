package main
import "fmt"
func cp(n int){
	if(n<2){
		fmt.Println("not prime")
	}
	for i:=2 ;i*i<n; i++{
		if n%i==0{
			fmt.Println("not prime")
			return
		}
	}
	fmt.Println("prime")
}
func main(){
	cp(5)
	cp(10)

}