package main
import (
	"net/http"
	"fmt"
	"io"
	"os"
)

func main() {
	// some errors http
	resp, err := http.DefaultClient.Get("www.baidu.com")
	if err != nil {
		fmt.Println("Error%s",err.Error)
		os.Exit(1)
	}
	defer resp.Body.close()
	io.Copy(os.Stdout, resp.Body)
}
