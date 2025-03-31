package main

import (
	"fmt"
	"time"
)

func main() {
	// sexyCount("nico")
	// sexyCount("flynn")

	// 매서드 앞 go 라는 명령어로 동시처리
	go sexyCount("nico")
	go sexyCount("flynn")
	time.Sleep(time.Second * 5)
}

func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)

	}

}
