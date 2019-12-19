package main

import (
	"fmt"
	"time"
)

func init() {
	fmt.Println("init will be before hello world")
}

func main() {
	fmt.Println("hello world")
	fmt.Println("today times:" + time.Now().String())
}
