package handlers

import "fmt"

func AuthPnlUser() {
	username, password := GetUserPnlCredentials()
	fmt.Println("Username:", username)
	fmt.Println("Password:", password)
}
