package main

import (
	"bst/bst"
	"fmt"
)

func main() {
	t := bst.NewTree(nil)
	/*

	   		    55
	   		  /   \
	   		 /     \
	   		/	    \
	   	   /         \
	   	  20          90
	   	/ 	 \       /
	   15	 50     80
	   		/      /
	   	   35     65


	*/
	t.Insert(55)
	t.Insert(20)
	t.Insert(90)
	t.Insert(80)
	t.Insert(50)
	t.Insert(35)
	t.Insert(15)
	t.Insert(65)
	fmt.Print("Inorder Traversal: ")
	bst.InorderTraversal(t.Root)
	fmt.Println()

	fmt.Print("Pre-order Traversal: ")
	bst.PreOrderTraversal(t.Root)
	fmt.Println()

	fmt.Print("Post-order Traversal: ")
	bst.PostOrderTraversal(t.Root)
	fmt.Println()
}
