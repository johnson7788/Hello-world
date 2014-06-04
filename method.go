package main
import "fmt"

type Integer int
func (a Integer) Les(b Integer) bool {
	return a<b
}
func main() {
	var a Integer =3
	if a.Les(5) {   //if a<b a.Les(2) is true, so Print..
		fmt.Println(a, "less 5" )
	} else {
		fmt.Println(a, "is bigger than 5")
	}
}
