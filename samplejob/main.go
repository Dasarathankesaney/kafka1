package main

import (
	"fmt"
)

func main() {

	for i := 1; i <= 100; i++ {
		//time.Sleep(time.Second * 1)
		fmt.Println("Processing invoice generation-", i)
	}

}
