//2.3.7 example
package main
import "fmt"
func modify(array [5]int) {
	array[0] = 10  //change the first element of array
	fmt.Println("In modify, Array values:", array)
}

func main() {
	array := [5]int{1,2,3,4,5}  //initial an array
	modify(array)     //try to invoking function
	fmt.Println("In main, Array values:", array)

	//Print every element
	for i,v := range array {
	fmt.Println("Array element", i, "=", v)
	}

//array slice, create an slice by created array

myArray := [10]int{1,2,3,4,5,6,7,8,9,10}
var mySlice []int = myArray[:5]

fmt.Println("myArray is:")
for _,v := range myArray {
	fmt.Print(v, " ")
}

fmt.Println("\n mySlice is:")
for _,v := range mySlice {
	fmt.Print(v, " ")
}
fmt.Println()

	//Use "make" create slice and , show the Slice cap() usage
	mySlice2 := make ([]int, 5, 10)   //5 is the length, 10 is the memory space, it is cap()
	fmt.Println("len=", len(mySlice2))
	fmt.Println("cap=", cap(mySlice2))
	mySlice2 = append(mySlice2,1,2,3)    //append some element in the end
	//slice copy
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}
	copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
	copy(slice1, slice2) // 只会复制slice2的3个元素到slice1的前3个位置





}
