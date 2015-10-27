package main 
import "fmt"

type Node struct{
	left *Node
	key int
	right *Node
}

func (this *Node) insert(data int){
	if data <this.key {
		if this.left!=nil{
			this.left.insert(data)
		}else{
			this.left=&Node{left:nil,key:data,right:nil}
		}
	}else{
		if this.right!=nil{
			this.right.insert(data)
		}else{
			this.right=&Node{left:nil,key:data,right:nil}
		}
	}
	
}

func printTree(root *Node) {
	if root.left!=nil{
		printTree(root.left)
	}
	fmt.Println(root.key)
	if root.right!=nil {
		printTree(root.right)
	}
	
}

func BST(n []int) *Node{
	bst:=new(Node)
	for i,val:=range n {
		if i==0 {
			bst.key=n[i]
		}else{
			bst.insert(val)
		}
	}
	return bst
}
func main() {
	nums:=[]int{2,8,1}
	root:=BST(nums)
	printTree(root)
	//fmt.Println(root.left)
}
