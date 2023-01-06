package main
import "fmt"

func main(){
	n:=4
	flag:=false
	if(n%2==0){
		flag=true
	}
	
	switch flag{
	case true:
		fmt.Println("even")
	case false:
		fmt.Println("odd")
	default:
		fmt.Println(flag)
	}

}