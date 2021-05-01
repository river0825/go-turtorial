package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

func walkRecursive(t *tree.Tree, ch chan int){

	if t.Left != nil {
		walkRecursive(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		walkRecursive(t.Right, ch)
	}
}
func Walk(t *tree.Tree, ch chan int) {
	walkRecursive(t, ch)
	close(ch)
}

func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for v1 := range ch1{
		v2 := <-ch2
		fmt.Println(v1, v2)
		if v1 != v2 {
			return false
		}
	}
	return true
}

//func main() {
//	b := Same(tree.New(1), tree.New(1))
//	fmt.Println(b)
//}
