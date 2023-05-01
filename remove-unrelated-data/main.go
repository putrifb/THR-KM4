package main

import "fmt"

func removeUnrelated(dataMap map[string]any, key string) map[string]any {
	//your code here
	delete(dataMap, key)

	return dataMap
}

func main() {
	data := map[string]any{"name": "Edo", "age": 20, "address": "jakarta"}
	fmt.Println(removeUnrelated(data, "age"))
}
