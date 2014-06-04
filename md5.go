package main

import "crypto/md5"
import "fmt"
import "io"
/*
func main() {
	h := md5.New()
	io.WriteString(h, "the fog is")
	fmt.Printf("%x\n", h.Sum(nil))
}
*/
func main() {
	h := "it is A"
	fmt.Println(md5.Sum(h))

}
