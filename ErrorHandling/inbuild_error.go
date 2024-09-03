package main

import (
	"fmt"
	"errors"
	"math/rand"
)

func inbuild_error(a, b int) (int, error){
	if b == 0 {
		return 0, errors.New("divide by Zero is not possible")
	}
	result := a/b
	return result, nil
}
func main() {
	for i := 0 ; i < 10 ; i++  {
		result , err := inbuild_error(rand.Intn(10), i)
		if err != nil {
			fmt.Println("Divide by Zeror Error")
		} else {
			fmt.Println(result)
		}
	}
}