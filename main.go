package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("Hello world from k8s...")

		time.Sleep(time.Second * 1)
	}
}
