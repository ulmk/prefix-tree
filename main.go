package main

import (
	"fmt"

	"github.com/ulmk/prefix-tree/ptree"
)

func main() {
	root := ptree.New()

	root.Insert("cat")
	root.Insert("cats")
	root.Insert("dog")

	root.Walk(func(n *ptree.Node) {
		// print each node
		fmt.Println(n)
	})

	fmt.Println("cat is word:", root.SubNodes['c'].SubNodes['a'].SubNodes['t'].Complete)

	other := ptree.New()
	other.Insert("car")
	other.Insert("card")

	root.Merge(other)

	fmt.Println("Merged tree")
	fmt.Println("car is word:", root.SubNodes['c'].SubNodes['a'].SubNodes['r'].Complete)
	fmt.Println("card is word:", root.SubNodes['c'].SubNodes['a'].SubNodes['r'].SubNodes['d'].Complete)

	root.Clear()

	fmt.Println("After clear")
	fmt.Println("cat is word:", root.SubNodes['c'].SubNodes['a'].SubNodes['t'].Complete)
}
