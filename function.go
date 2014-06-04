package main
import "fmt"
import "errors"
//函数名是 Compute,参数列表是 a,b, 返回值是 result,err, 返回语句是 return
func Compute(a int, b int)(result float64, err error) {
	if a<b {
	err= errors.New("Should be non-negative numbers")
	}
	result = 5.1
	return
}
func main(){
	fmt.Println(Compute(1,2))
}
