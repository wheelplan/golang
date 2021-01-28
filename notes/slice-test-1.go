package main

import "fmt"

func main() {

	data := []string{"", "pink", "red", "yellow", "", "blue", ""}
	afterData := noEmpty(data)

	fmt.Print(afterData, len(afterData), cap(afterData))
}

func noEmpty(data []string) []string {

	var newData []string

	for _, v := range data {
		if v != "" {
			newData = append(newData, v)
		}
	}

	return newData
}
