package main

import (
	"fmt"
	"os/user"
	"time"
)

//func init() {
//	fmt.Println("init")
//}

func bazz() {
	fmt.Println("bazz")
}

func main() {
	bazz()
	fmt.Println("Hello World", time.Now())
	fmt.Println(user.Current())
}
