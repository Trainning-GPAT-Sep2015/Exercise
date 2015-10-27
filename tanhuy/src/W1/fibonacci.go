package main 
import "fmt"

func fibo(n int) int {
	if (n<0){
		panic("Invalid input")
	}else if(n==0 || n==1){
		return n
	}else {
		return fibo(n-1)+ fibo(n-2)
	}
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(fibo(i))
	}
}
