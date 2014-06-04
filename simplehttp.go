package main
import (
	"net"
	"os"
	"bytes"
	"fmt"
	"io"
)

func main() {
	arg := os.Args[1:]
	if len(arg) != 1 {
		fmt.Println("USAGE: host, Must be run by root user")
		os.Exit(1)
	}
	service := arg[0]
	conn, err := net.Dial("tcp", service)
	checkError(err)

	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)

	result, err := readFully(conn)
	checkError(err)

	fmt.Println(string(result))
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Println(os.Stderr,"Fatal error:%s", err.Error())
		os.Exit(1)
	}
}
func readFully(conn net.Conn) ([]byte, error) {
	defer conn.Close()
	result := bytes.NewBuffer(nil)
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF{
				break
			}
		return nil, err
		}
	}
	return result.Bytes(), nil
}

