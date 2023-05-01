package main

import "fmt"

func AddElement(data []int, newData int, position string) []int {
	//your code here!
	if position == "up" {
		return append([]int{newData}, data...)
	} else if position == "down" {
		return append(data, newData)
	}
	return nil
}

func main() {
	h := []int{1, 2, 3, 4, 5}
	n := 6
	g := "up"
	fmt.Println(AddElement(h, n, g))

	a := []int{1, 2, 3, 4, 5}
	c := 6
	d := "down"
	fmt.Println(AddElement(a, c, d))
}
