package main
import (
	"io"
	"log"
	"net/http"
)
// Please access browser to http://10.13.2.36:8080/hello
func helloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w,"Hello,world!")    //this is content
}
func main() {
	http.HandleFunc("/hello", helloHandler)   //hello is the path
	err :=http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err.Error())
	}
}
