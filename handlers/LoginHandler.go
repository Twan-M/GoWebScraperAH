package handlers

import (
	"log"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/cookies"
)

func LoginHandler() {
	// Instantiate a new Colly collector
	c := colly.NewCollector()

	cookieJar, err := cookies.NewCollector(
		cookies.AllowedDomains("example.com"),
	)
	if err != nil {
		panic(err)
	}
	c.SetCookieJar(cookieJar)

	// Define the login URL, credentials and the POST data
	loginURL := "https://euidp.aholddelhaize.com/"
	username := "pnl05350"
	password := "Patat123"
	loginCredentials := map[string]string{
		"username": username,
		"password": password,
	}

	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36 Edg/114.0.0.0"

	// Use the OnHTML function to log in to the website
	c.OnHTML("form", func(e *colly.HTMLElement) {
		// Submit the login form with the POST data
		err := c.Post(loginURL, loginCredentials)
		if err != nil {
			log.Fatal(err)
		}
	})

	// Use the OnResponse function to verify if the login was successful
	c.OnResponse(func(r *colly.Response) {
		if r.StatusCode == 200 {
			// If the login was successful, access the protected page
			c.Visit("https://sam.ahold.com/etm/time/timesheet/etmTnsMonth.jsp")
			log.Println("Login successful with code: ", r.StatusCode, " and body: ", string(r.Body))
		} else {
			log.Println("Login failed with code: ", r.StatusCode, " and body: ", string(r.Body))
		}
	})

	// Start the collector by visiting the login page
	c.Visit("https://euidp.aholddelhaize.com/")
}
