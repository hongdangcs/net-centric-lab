package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"netcentric/lab3"
	"netcentric/lab4"
	"strings"
	"time"
)

var hangmanPlayers = make(map[net.Conn]int)
var hangmanWord string
var currentHangmanPlayer net.Conn

func main() {

	hangmanWord = lab4.GetRandomWord()
	hangmanWord = strings.ToLower(hangmanWord)
	censoredWord := ""
	for range hangmanWord {
		censoredWord += "_"
	}
	timer := time.NewTimer(60 * 60 * 60 * time.Second)

	file, err := ioutil.ReadFile("lab3/users.json")
	if err != nil {
		log.Fatalf("failed to read file: %s", err)
	}

	var users lab3.Users
	err = json.Unmarshal(file, &users)
	if err != nil {
		log.Fatalf("failed to unmarshal JSON: %s", err)
	}

	userMap := make(map[string]lab3.User)
	for _, user := range users.Users {
		userMap[user.Username] = user
	}

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go lab3.HandleConnection(conn, userMap, &hangmanPlayers, &hangmanWord, &currentHangmanPlayer, &censoredWord, timer)
	}

}
