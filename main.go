package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Start!!!")

	ch := make(chan bool)

	go func() {
		name := "sub"
		for index := 0; index < 10; index++ {
			<-ch
			fmt.Printf("%s >", name)
			fmt.Printf("I'm %s routine...(count %d)\n", name, index)
			time.Sleep(time.Second)
		}
	}()

	name := "main"
	for index := 0; index < 10; index++ {
		fmt.Printf("%s >", name)
		fmt.Printf("I'm %s routine...(count %d)\n", name, index)
		ch <- true
		time.Sleep(time.Second)
	}
}
