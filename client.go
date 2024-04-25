package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to the server: ", err)
		os.Exit(1)
	}
	defer conn.Close()

	token := ""

	go func() {
		reader := bufio.NewReader(conn)
		for {
			message, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading from server:", err)
				break
			}
			fmt.Println("Server response: " + message)

			if strings.Contains(message, "Your token is") {
				token = strings.Split(message, ": ")[1]
				token = strings.TrimSpace(token)
			}
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		text := scanner.Text()
		if token != "" {
			text = token + "_" + text
		}
		fmt.Fprintf(conn, text+"\n")
	}
}
