// execute ## go run main.go
package main // 패키지 이름 , main.go 컴파일 원하면 필요함
import "fmt"

func main() {
	// const name string = "jaewon"
	// name = "Suz" can't change

	// 축약코드 사용시 Go가 타입을 찾아줌
	name := "jaewon" // == var name string = "jaewon"
	name = "하이"      // can change
	fmt.Println(name)
}
