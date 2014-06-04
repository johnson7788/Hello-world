package main
import "fmt"

//PersonInfo include a person information
type PersonInfo struct {
	ID string
	Name string
	Address string
}
func main() {
var	personDB map[string] PersonInfo
	personDB = make(map[string] PersonInfo)

//insert some data
	personDB["12"] = PersonInfo{"128391","Tom", "Room 22"}
	personDB["00"] = PersonInfo{"000111", "Jack", "Room 11"}
	person, ok := personDB["12"]  //find a map key
if ok {
	fmt.Println("ID128391 person is:", person.Name)
} else{
	fmt.Println("Did not find this person")
}
	//delete (personDB, "12")

}
