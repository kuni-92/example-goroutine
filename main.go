package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Start!!!")

	go func() {
		for index := 0; index < 10; index++ {
			fmt.Printf("I'm sub routine...(count %d)\n", index)
			time.Sleep(2 * time.Second)
		}
	}()
	for index := 0; index < 10; index++ {
		fmt.Printf("I'm main routine...(count %d)\n", index)
		time.Sleep(1 * time.Second)
	}
}