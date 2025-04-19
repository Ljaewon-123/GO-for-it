package main

import "syscall/js"

func main() {
	js.Global().Call("alert", js.ValueOf("Hello from syscall/js!"))
}
