package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}

// go build -gcflags="-S" main.go 2> output.asm
// go tool objdump -s "main.main" ./main > output.asm

// GOARCH=amd64 GOOS=windows go build main.go
// GOARCH=arm64 GOOS=darwin go build main.go
// GOARCH=arm64 GOOS=linux go build main.go
