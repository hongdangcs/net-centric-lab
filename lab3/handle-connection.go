package lab3

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net"
	"netcentric/lab4"
	"os"
	"strconv"
	"strings"
	"time"
)

func HandleConnection(conn net.Conn, users map[string]User, hangmanPlayers *map[net.Conn]int, hangmanWord *string, currentHangmanPlayer *net.Conn, censoredWord *string, timer *time.Timer) {
	defer conn.Close()

	fmt.Println("New connection from", conn.RemoteAddr())
	ranNumber := 0

	for {
		netData, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Println(err)
			return
		}
		temp := strings.TrimSpace(string(netData))
		if strings.HasPrefix(temp, "login") {
			paths := strings.Split(temp, " ")
			if len(paths) != 3 {
				fmt.Fprintf(conn, "Invalid login format\n")
				continue
			}
			username := paths[1]
			password := paths[2]
			user, exists := users[username]
			if exists && user.Password == EncryptPassword(password) {
				token := GenerateUniqueToken(username)
				tokenStr := strconv.Itoa(token)
				tokens[tokenStr] = username
				fmt.Fprintf(conn, "Your token is: %s\n", tokenStr)
			} else {
				fmt.Fprintf(conn, "Invalid user\n")
			}
			continue
		}

		if strings.Contains(temp, "download") {
			paths := strings.Split(temp, "_")
			token := paths[0]
			username := TokenCheck(token, tokens)
			if username == "" {
				fmt.Fprintf(conn, "Invalid token\n")
				continue
			}

			fileName := strings.Split(temp, " ")[1]
			file, err := os.ReadFile(fileName)
			if err != nil {
				fmt.Fprintf(conn, "File not found\n")
				continue
			}
			fileContent := string(file)
			fmt.Fprintf(conn, fileContent)
			continue
		}

		if strings.Contains(temp, "hangman") {
			paths := strings.Split(temp, "_")
			token := paths[0]
			username := TokenCheck(token, tokens)
			if username == "" {
				fmt.Fprintf(conn, "Invalid token\n")
				continue
			}
			(*hangmanPlayers)[conn] = 0
			if len(*hangmanPlayers) < 2 {
				fmt.Fprintf(conn, "Waiting for another player\n")
				continue
			}

			*currentHangmanPlayer = conn

			for player := range *hangmanPlayers {
				fmt.Fprintf(player, "Game started, guess the word: %s\n", *censoredWord)
				if player == conn {
					fmt.Fprintf(player, "You are the guesser\n")
					timeCountDown(hangmanPlayers, currentHangmanPlayer, timer)

				} else {
					fmt.Fprintf(player, "Let wait for the other guess\n")
				}
			}

			continue
		}

		if strings.Contains(temp, "guess") {
			paths := strings.Split(temp, "_")
			token := paths[0]
			username := TokenCheck(token, tokens)
			if username == "" {
				fmt.Fprintf(conn, "Invalid token\n")
				continue
			}
			if conn != *currentHangmanPlayer {
				fmt.Fprintf(conn, "Not your turn\n")
				continue
			}
			guess := paths[1]
			if len(strings.Split(guess, " ")) > 1 {
				guess = strings.Split(guess, " ")[1]
			}
			if len(guess) > 1 {
				fmt.Fprintf(conn, "Invalid guess\n")
				continue
			}
			guess = strings.ToLower(guess)
			if strings.Contains(*hangmanWord, guess) {
				for i, c := range *hangmanWord {
					if string(c) == guess {
						*censoredWord = (*censoredWord)[:i] + guess + (*censoredWord)[i+1:]
						(*hangmanPlayers)[conn] += 10
					}

				}
				if *censoredWord == *hangmanWord {
					highestScore := 0
					for player := range *hangmanPlayers {
						if (*hangmanPlayers)[player] > highestScore {
							highestScore = (*hangmanPlayers)[player]
						}
					}
					for player := range *hangmanPlayers {
						fmt.Fprintf(player, "Game over, the word is: %s\n", *hangmanWord)
						if (*hangmanPlayers)[player] == highestScore {
							fmt.Fprintf(player, "You win with score: %d\n", highestScore)
						} else {
							fmt.Fprintf(player, "You lose with score: %d\n", (*hangmanPlayers)[player])
						}
					}
					*hangmanWord = lab4.GetRandomWord()
					*hangmanWord = strings.ToLower(*hangmanWord)
					*censoredWord = ""
					for range *hangmanWord {
						*censoredWord += "_"
					}
					timeCountDown(hangmanPlayers, currentHangmanPlayer, timer)
				} else {
					fmt.Fprintf(conn, "Correct guess, let continue\n")
					for player := range *hangmanPlayers {
						if player != conn {
							fmt.Fprintf(player, "Word: %s\n", *censoredWord)
						}
					}
					// reset timer
					(*timer).Reset(30 * time.Second)
				}
			} else {
				fmt.Fprintf(conn, "Incorrect guess, you lose your turn\n")
				switchPlayer(hangmanPlayers, currentHangmanPlayer, timer)
				fmt.Fprintf(*currentHangmanPlayer, "Your turn\n")

			}
			for player := range *hangmanPlayers {
				fmt.Fprintf(player, "Guesser: %s\n", guess)
				fmt.Fprintf(player, "Word: %s\n", *censoredWord)
			}
			continue
		}

		if strings.Contains(temp, "start") {
			ranNumber = rand.Intn(100)
			fmt.Println("Random number: ", ranNumber)
			fmt.Fprintf(conn, "Guess a number between 0 and 100\n")
			continue
		}

		if strings.Contains(temp, "_") {
			if ranNumber == 0 {
				fmt.Fprintf(conn, "No game started, type 'start' to start a game\n")
				continue
			}
			paths := strings.Split(temp, "_")
			token := paths[0]
			username := TokenCheck(token, tokens)
			if username == "" {
				fmt.Fprintf(conn, "Invalid token\n")
				continue
			}
			guess, _ := strconv.Atoi(paths[1])
			if guess == ranNumber {
				fmt.Fprintf(conn, "Correct\n")
			} else if guess < ranNumber {
				fmt.Fprintf(conn, "Too low\n")
			} else {
				fmt.Fprintf(conn, "Too high\n")
			}
			continue
		}

		fmt.Println("Received: ", temp)
		result := temp
		fmt.Fprintf(conn, "%s\n", result)
	}
}

func switchPlayer(hangmanPlayers *map[net.Conn]int, currentHangmanPlayer *net.Conn, timer *time.Timer) {
	current := false
	firstPlayer := new(net.Conn)
	for player := range *hangmanPlayers {
		if *firstPlayer == nil {
			*firstPlayer = player
		}
		if player == *currentHangmanPlayer {
			current = true
			*currentHangmanPlayer = nil
			continue
		}
		if current {
			*currentHangmanPlayer = player
		}

	}
	if *currentHangmanPlayer == nil {
		*currentHangmanPlayer = *firstPlayer
	}
	timeCountDown(hangmanPlayers, currentHangmanPlayer, timer)
}

func timeCountDown(hangmanPlayers *map[net.Conn]int, currentHangmanPlayer *net.Conn, timer *time.Timer) {
	timer.Reset(30 * time.Second)
	go func() {
		<-timer.C
		if _, ok := (*hangmanPlayers)[*currentHangmanPlayer]; ok {
			fmt.Fprintf(*currentHangmanPlayer, "Time's up! Next player's turn.\n")
			switchPlayer(hangmanPlayers, currentHangmanPlayer, timer)
			fmt.Fprintf(*currentHangmanPlayer, "Your turn\n")
		}
	}()
}
