package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

func ShowHomePage() {
	// Set the protected page URL
	pageURL := "https://sam.ahold.com/az_ahsam_jct/home.htm"

	// Retrieve Cookies
	cookies, err := GetCookie()
	if err != nil {
		log.Fatal("Error getting Cookies")
	}

	// Create a new GET request with the cookies set in the headers
	req, err := http.NewRequest("GET", pageURL, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Add the cookies obtained from the login process to the request headers
	for _, cookie := range cookies {
		req.AddCookie(cookie)
		fmt.Println("Cookie:")
		fmt.Println(cookie)
	}

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Request failed with status code:", resp.StatusCode)
		return
	}

	// Read the response body and print the page title
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	re := regexp.MustCompile("<title>(.*?)</title>")
	title := re.FindStringSubmatch(string(body))[1]
	fmt.Println(title)
}
