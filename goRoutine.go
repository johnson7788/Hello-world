package main
import "fmt"
import "time"

func GetSum(num int, divider int, result chan int) {
	sum := 0
	for value := 0;  value<num; value++ {
		if value%divider == 0 {
			sum +=value
		}
	}
	result <- sum
}

func main() {
	Limit := 100
	result := make(chan int, 3)
	tStart := time.Now()
	go GetSum(Limit,3,result)
	go GetSum(Limit,5,result)
	go GetSum(Limit,15,result)

	sum3,sum5,sum15 := <-result, <-result, <-result
	sum := sum3+sum5-sum15
	tEnd := time.Now()
	fmt.Println(sum)
	fmt.Println(tEnd.Sub(tStart))

}
