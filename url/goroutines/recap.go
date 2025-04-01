package goroutines

import (
	"fmt"
	"time"
)

func Recap() {
	c := make(chan string)
	people := [5]string{"nico", "flynn", "dal", "japanguy", "larry"}
	for _, person := range people {
		go isSexy(person, c)
	}
	for i := 0; i < len(people); i++ {
		fmt.Println(<-c)
	}

	// 모든 채널이 지나간후에 받음 마치... await처럼
	fmt.Println("When RUN?")

}

func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 10)
	c <- person + " is sexy"
}
