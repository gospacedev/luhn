package main

import (
	"fmt"
	"github.com/gocrazygt/luhn"
)

func main() {
	a := luhn.Check("79927398713")
	b := luhn.Check("1111")
	fmt.Println(a)
	fmt.Println(b)
}
