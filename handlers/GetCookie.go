package handlers

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
)

func GetCookie() ([]*http.Cookie, error) {
	// Set the login form URL
	loginURL := "https://euidp.aholddelhaize.com/pkmslogin.form"

	// Set the login credentials
	username, password := GetUserPnlCredentials()
	fmt.Println(username + password)

	// Prepare the form data
	formData := url.Values{}
	formData.Set("username", username)
	formData.Set("password", password)

	// Encode the form data
	formEncoded := formData.Encode()

	// Create a new POST request with the login form data
	req, err := http.NewRequest("POST", loginURL, bytes.NewBufferString(formEncoded))
	if err != nil {
		fmt.Println("Error creating request:", err)
	}

	// Set the Content-Type header to application/x-www-form-urlencoded
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
	}

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Login failed with status code:", resp.StatusCode)
	}

	// Extract the cookies set by the server
	var cookies []*http.Cookie
	cookies = resp.Cookies()

	return cookies, err
}
