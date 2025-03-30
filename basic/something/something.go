package something

import "fmt"

// none export is private
func sayBye() {
	fmt.Println("Bye")
}

// export use another files
func SayHello() {
	fmt.Println("Hello")
}
