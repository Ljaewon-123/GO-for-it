package channels

import (
	"fmt"
	"time"
)

func Channels() {
	c := make(chan bool) // channel
	people := [2]string{"nico", "flynn"}
	for _, person := range people {
		go isSexy(person, c)
	}
	fmt.Println(<-c) // <-c 채널로 들어올때까지 기다리는 거임
	fmt.Println(<-c)
	fmt.Println(<-c) // error dead lock
}

func isSexy(person string, c chan bool) {
	time.Sleep(time.Second * 5)
	fmt.Println(person)
	c <- true
}
