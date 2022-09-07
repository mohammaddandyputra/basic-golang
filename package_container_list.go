package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()

	data.PushBack("Mohammad")
	data.PushBack("Dandy")
	data.PushBack("Putra")
	data.PushFront("Mr.")

	fmt.Println(data.Front().Prev())

	fmt.Println(data.Front().Value)
	fmt.Println(data.Back().Value)
}
