package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Start!!!")

	quit := make(chan bool)
	go func() {
		for index := 0; index < 10; index++ {
			fmt.Printf("I'm sub routine...(count %d)\n", index)
			time.Sleep(2 * time.Second)
		}
		quit <- true
	}()
	for index := 0; index < 10; index++ {
		fmt.Printf("I'm main routine...(count %d)\n", index)
		time.Sleep(1 * time.Second)
	}
	<-quit
}
