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

// 알아서 return변수로 return해줌
func lanAndUpperNaked(name string) (length int, uppercase string) {
	// 이미 선언된 return변수들을 사용
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

// func가 리턴한 이후 실행할 동작을 정의
func lanAndUpperDefer(name string) (length int, uppercase string) {
	// func값을 return하고 실행
	defer fmt.Println("I'm done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func superAdd(numbers ...int) int {
	total := 0
	for index, number := range numbers {
		fmt.Println("index: ", index, "number: ", number)
		total += number
	}
	return total
}

func canIDrink(age int) bool {
	// if안에 변수만들기 가능   ; 구분자
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true
}

func canIDrinkSwitch(age int) bool {
	// switch age {
	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 18:
		return true
	}
	return false
}

func Pointers() {
	a := 2
	b := &a
	// a = 5
	*b = 20            // b를 이용해서 같은 메모리 주소의 a를 변경가능 (b는 a를 보는 포인터 이기에 가능)
	fmt.Println(a, *b) // &a, *a   메모리 주소, 주소값
}

func main() {
	fmt.Println(multiply(2, 2))

	totalLen, upperName := lanAndUpper("jaewon")
	fmt.Println(totalLen, upperName)

	totalLen2, _ := lanAndUpper("jaewon")
	fmt.Println(totalLen2)

	repeatMe("jaewon", "suz", "dal", "marl")

	fmt.Println(lanAndUpperNaked("jaewon")) // 6 JAEWON

	// I'm done
	// 6 JAEWON
	fmt.Println(lanAndUpperDefer("jaewon"))

	result := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(result)

	fmt.Println(canIDrink(16))
	fmt.Println(canIDrinkSwitch(18))

	fmt.Println("Pointer!!")

	Pointers()
}
