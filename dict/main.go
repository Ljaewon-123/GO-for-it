package main

import (
	"fmt"

	"github.com/Ljaewon-123/GO-for-it/dict/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	word := "hello"
	dictionary.Add(word, "First")
	err := dictionary.Update(word, "Second")
	if err != nil {
		fmt.Println(err)
	}
	result, _ := fmt.Println(dictionary.Search(word))
	fmt.Println(result)
}
