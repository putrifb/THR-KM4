package main

import "fmt"

func removeLeftRight(arr []any, position string) []any {
	//yout code here!
	if position == "left" {
		return arr[1:]
	} else if position == "right" {
		return arr[:len(arr)-1]
	}

	return nil
}

func main() {
	fmt.Println(removeLeftRight([]any{1, 2, 3, 4, 5}, "left"))
	fmt.Println(removeLeftRight([]any{1, 2, 3, 4, 5}, "right"))
}
