package main

import "base91-go/base91"

func main() {
	encoded := base91.EncodeToString([]byte("hello world"))
	println(encoded)

	decoded := base91.DecodeString(encoded)
	println(string(decoded))
}
