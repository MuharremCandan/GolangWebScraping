package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gocolly/colly"
)

func main() {
	fName := "data.csv"

	//create an excel file to save our datas
	file, err := os.Create(fName)

	//check if there is an error while creating the file
	if err != nil {
		log.Fatalf("Couldn't create file , err : %q", err)
		return
	}

	//close the file at the end of the method
	defer file.Close()

	//  write on file our datas
	writer := csv.NewWriter(file)

	defer writer.Flush()

	// data collector from given link in it
	c := colly.NewCollector(
		colly.AllowedDomains("internshala.com"),
	)

	// data pulling  among html codes
	c.OnHTML(".internship_meta", func(h *colly.HTMLElement) {

		//writing datas in the file
		writer.Write([]string{
			h.ChildText("a"),
			h.ChildText("span"),
		})
	})

	// pulling all  datas from other pages
	for i := 0; i < 312; i++ {
		fmt.Printf("Scraping Page : %d\n", i)
		c.Visit("https://internshala.com/internships/page-" + strconv.Itoa(i))
		log.Printf("Scraping Complete\n")
		log.Println(c)
	}

}
