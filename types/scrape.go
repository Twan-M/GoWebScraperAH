package types

import "gorm.io/gorm"

type Scrape struct {
	gorm.Model
	Text string `json:"text"`
}
