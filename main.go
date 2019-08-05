package main

import (
	"fmt"

	"golang.org/x/sys/unix"
)

func main() {
	fmt.Println(unix.Getgid())
}
