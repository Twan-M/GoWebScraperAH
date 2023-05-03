package handlers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetUserHandler() (string, string) {
	var pnlUsername string
	var pnlPassword string

	fmt.Print("Enter PNL Username: ")
	fmt.Scanln(&pnlUsername)
	fmt.Print("Enter PNL Password: ")
	bytePassword, err := bufio.NewReader(os.Stdin).ReadBytes('\n')
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	pnlPassword = strings.TrimSpace(string(bytePassword))

	return pnlUsername, pnlPassword
}
