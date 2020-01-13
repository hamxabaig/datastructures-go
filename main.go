package main

import linked "github.com/hamxabaig/datastructures-go/data_structures/linked_list"

import "fmt"

func main() {
	var list = linked.New()
	list.Prepend(1)
	list.Prepend(2)
	list.Append(3)
	list.Append(4)
	list.Print()
	fmt.Println(list.Length())
	list.Append(3)
	fmt.Println(list.Length())
}
