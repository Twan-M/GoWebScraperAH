package repositories

import "GoScraper/types"

func AddTitle(title types.Scrape) error {
	if err := connection().Create(&title).Error; err != nil {
		return err
	}
	return nil

}
