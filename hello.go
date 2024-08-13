package myhello

import (
	"fmt"
	"time"
)

func PrintHello() {
	fmt.Println("hello", time.Now().Unix())
}
