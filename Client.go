package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	conn, err := net.DialTimeout("tcp", "localhost:8080", 5*time.Second) // Dial with timeout
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

		// Set write deadline
		err = conn.SetWriteDeadline(time.Now().Add(5 * time.Second))
		if err != nil {
			fmt.Println("Error setting write deadline:", err)
			return
		}
	}
}
