package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "hello world hello world"

	strs := strings.Split(s, " ")
	ss := strings.Join(strs, " ")
	fmt.Println(strs, ss) //4

}
