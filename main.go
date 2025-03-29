// execute ## go run main.go
package main // 패키지 이름 , main.go 컴파일 원하면 필요함
import (
	"fmt"
	"strings"
)

//	func multiply(a int, b int) int {
//		return a * b
//	}
func multiply(a, b int) int {
	return a * b
}

// multi return function
func lanAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	fmt.Println(multiply(2, 2))

	totalLen, upperName := lanAndUpper("jaewon")
	fmt.Println(totalLen, upperName)

	totalLen2, _ := lanAndUpper("jaewon")
	fmt.Println(totalLen2)

	repeatMe("jaewon", "suz", "dal", "marl")
}
