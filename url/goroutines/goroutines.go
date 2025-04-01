package goroutines

import (
	"fmt"
	"time"
)

func GoroutinesWithCount() {
	// sexyCount("nico")
	// sexyCount("flynn")

	// 매서드 앞 go 라는 명령어로 동시처리
	go sexyCount("nico")
	go sexyCount("flynn")

	// goroutines은 main함수가 실행중인 동안만 유효해서 이게 없으면 종료되어버림
	time.Sleep(time.Second * 5)
}

func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)

	}

}
