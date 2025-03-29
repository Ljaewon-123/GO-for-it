// execute ## go run main.go
package main // 패키지 이름 , main.go 컴파일 원하면 필요함
import (
	"fmt"

	"github.com/Ljaewon-123/GO-for-it/something"
)

func main() {
	fmt.Println("Hello world!")
	something.SayHello()

}
