package main

import (
	"database/sql"
	"fmt"
	"net/url"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gocolly/colly"
)

func getData(url *url.URL) {

	collection := colly.NewCollector()

	//////

	collection.OnHTML(".amount--3NTpl", func(e *colly.HTMLElement) {
		price := e.Attr("class")
		fmt.Println(e.Text)
		collection.Visit(e.Request.AbsoluteURL(price))
	})
	collection.OnHTML(".word-break--2nyVq", func(e *colly.HTMLElement) {
		information := e.Attr("class")
		fmt.Println(e.Text)
		collection.Visit(e.Request.AbsoluteURL(information))
	})

	collection.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	collection.Visit(url.String())

}

func main() {
	fmt.Print("Please enter here catagory Type ->>> ")
	var category string
	fmt.Scanln(&category)

	fmt.Print("Please enter here District ->>> ")
	var District string
	fmt.Scanln(&District)

	collectionn := colly.NewCollector()

	dataBase, _ := sql.Open("mysql", "root:ijse@tcp(127.0.0.1:3306)/ikman")

	collectionn.OnHTML(".gtm-normal-ad", func(element *colly.HTMLElement) {
		model := element.ChildText(".heading--2eONR")
		descr := element.ChildText(".description--2-ez3")
		price := element.ChildText(".price--3SnqI")
		fmt.Println("\n")
		fmt.Println("\tmodel : ", model)
		fmt.Println("\tprice : ", price)
		fmt.Println("\tdescr : ", descr)

		insert, err := dataBase.Query("INSERT INTO ikman (category, District, model, price, descr) VALUES (?, ?, ?, ?, ?)", category, District, model, price, descr)
		check(err)
		defer insert.Close()
	})

	collectionn.OnRequest(func(t *colly.Request) {
		url := (t.URL)
		getData(url)

	})

	_ = collectionn.Visit("https://ikman.lk/en/ads/" + District + "/" + category)

}

func check(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
