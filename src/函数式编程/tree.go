package main

import (
	"fmt"
	"reflect"
)

type Node struct{
	value int
	left *Node
	right *Node
}

func (node *Node)Print(){
	fmt.Println(node.value)
}

func (node *Node) Travel(){
	if node==nil{
		return
	}

	node.left.Travel()
	node.Print()
	node.right.Travel()
}

//建立一棵树
func MakeTree()*Node{
	n1:=Node{3,nil,nil}
	n2:=Node{0,nil,nil}
	n3:=Node{5,nil,nil}
	n4:=Node{2,nil,nil}
	n5:=Node{4,nil,nil}

	n1.left=&n2
	n1.right=&n3
	n2.right=&n4
	n3.left=&n5

	return &n1
}
/*-------函数式编程便利tree-----------------*/
func (node *Node) Traverse(){

	node.TraverseFunc(func(node1 *Node){//穿个函数打印
		node1.Print()
	})
}

func (node *Node) TraverseFunc(f func(node *Node)){
	if node==nil{
		return
	}

	node.left.TraverseFunc(f)
	f(node)
	node.right.TraverseFunc(f)
}




func main() {
	//head:=MakeTree()
	//head.Traverse()

	//数节点
	/*count:=0
	head.TraverseFunc(func(node *Node){
		count++
	})

	fmt.Printf("the count of node is : %d",count)*/
	a:=map[string]interface{}{
		"name":"xishilong",
		"age":20,
	}

	c,ok:=a["name"].(string)
	fmt.Println(c)
	fmt.Println(reflect.TypeOf(ok))
}

