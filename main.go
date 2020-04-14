package main

import (
	"fmt"
	"time"

	"gose/internal/scraper"
)

func main() {
	var keywords = []string{"Arham Abiyan"}
	for _, keyword := range keywords {
		res, _ := scraper.GoogleScrape(keyword, "id", "id")
		fmt.Println("-----------------------")
		fmt.Println(keyword)
		for _, item := range res {
			fmt.Println(item)
		}
		time.Sleep(time.Second * 5)
	}
}
