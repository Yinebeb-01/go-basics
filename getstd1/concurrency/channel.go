package main

import "fmt"

func channel(name chan string) {
	var a chan int // empty/nil channel which has no any use,
	fmt.Println("value of empty channnel:", a)
	mychan := make(chan int) // this chan, formed via make comes useful to pass data

	fmt.Println("value of make channel: ", mychan)

	fmt.Println("hello" + <-name) // passing data from channel to stdout.
	//let we see effect of closing a channel
	<-name
}

// square - more on channel
func square(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i * i //sending square of i to the channel, to retrieve later via a goroutine
		fmt.Println("from go routine")
	}
	close(c)
}


