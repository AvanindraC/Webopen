package main

import (
	"fmt"
	"os/exec"
)

func main() {
	var a string
	fmt.Println("Enter the url")
	fmt.Scanln(&a)
	exec.Command("rundll32", "url.dll,FileProtocolHandler", a).Start()
}
