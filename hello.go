package main

import (
	"fmt"
	"github.com/shibas11/go-hello-world/stringutil"
)

func main() {
	s := "Hello, world."
	fmt.Println(s)
	fmt.Println(stringutil.Reverse(s))
}
