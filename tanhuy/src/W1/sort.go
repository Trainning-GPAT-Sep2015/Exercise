package main 
import "fmt"

func sort(n []int) {
	swapped:=true
	for swapped {
		swapped=false;
		for i := 0; i < len(n)-1; i++ {
			if n[i]<n[i+1]{
				swap(n,i,i+1)
				swapped=true
			}
		}
	}
}

func swap(arr []int,a,b int){
	c:=arr[a]
	arr[a]=arr[b]
	arr[b]=c
}

func main() {
	num:=[]int{5,3,25,1}
	fmt.Println("Unsort Array: " ,num)
	sort(num)
	fmt.Println("Sorted Array: ", num)
}
