//任意类型的不定参数 2.4
package main
import "fmt"
func Myprint(args ...interface{}) {
	for _,v := range args {
		switch v.(type) {
			case int:
				fmt.Println(v, "is a int")
			case string:
				fmt.Println(v, "is a string")
			case int64:
				fmt.Println(v, "is a int64")
			default:
				fmt.Println(v, "is unknow type")
		}
	}
}
func main() {
	v1 :=1
	v2 :=234
	v3 :="Hi"
	v4 :=3.23
	var v5 int64 = 666
	Myprint(v1,v2,v3,v4,v5)

}
