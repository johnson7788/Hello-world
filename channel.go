package main
import "fmt"

func Count(ch chan int) {
//我们通过ch <- 1语 句向对应的channel中写入一个数据。在这个channel被读取前,这个操作是阻塞的
	ch <-1
	fmt.Println("Counting")
}

func main() {
	chs := make([]chan int, 10)
	for i:=0; i<10; i++ {
		chs[i] =make(chan int)
		go Count(chs[i])
	}
	for _,ch := range(chs) {
		<-ch
	}
}
