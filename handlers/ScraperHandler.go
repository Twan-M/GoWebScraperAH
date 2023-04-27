package handlers

import (
	"GoScraper/repositories"
	"GoScraper/types"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

func TitleScraper() {
	url := "https://euidp.aholddelhaize.com/"

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	re := regexp.MustCompile("<title>(.*?)</title>")
	title := re.FindStringSubmatch(string(body))[1]

	s := types.Scrape{Text: title}
	err = repositories.AddTitle(s)
	if err != nil {
		log.Panic("Error adding Title in DB: " + err.Error())
	}

	fmt.Println(title)
}
