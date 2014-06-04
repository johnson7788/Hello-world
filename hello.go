/* hello.go - My first Golang program */
package main

import "fmt"
import "os"

func main() {
	fmt.Printf("Hello, world\n")
	argsnew := os.Args[1:]
	for i, v := range argsnew {
		fmt.Println(i, " ", v)
	}
}
