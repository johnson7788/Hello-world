package main
import "fmt"
import "errors"

func main() {
	err := errors.New("I am difine the erros again")
	if err != nil {
		fmt.Println(err)
	}
}
