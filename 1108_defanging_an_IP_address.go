package main

import (
	"fmt"
	"strings"
)

func defangIPaddr(address string) string {
	return strings.ReplaceAll(address, ".", "[.]")
}

func main() {
	address := "1.1.1.1"
	output := defangIPaddr(address)
	fmt.Println(output)
}
