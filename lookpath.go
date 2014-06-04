package main
import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	path, err := exec.LookPath("ping")
	if err != nil {
		log.Fatal("Error,Don't find the command in $PATH")
	}
	fmt.Printf("command is at %s\n", path)
}
