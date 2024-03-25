package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()
	fmt.Println("Connected to server")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter text: ")
		scanner.Scan()
		text := scanner.Text()
		if text == "exit" {
			break
		}
		_, err := conn.Write([]byte(text))
		if err != nil {
			fmt.Println("Error sending data:", err)
			return
		}
	}
}
