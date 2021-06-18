package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	htmlquery "github.com/antchfx/htmlquery"
)

func getPriceFromGoogle() (float64, error) {
	url := "https://www.google.com/search?q=usd+to+cop&oq=usd+to+cop&aqs=chrome..69i57.1749j0j4&sourceid=chrome&ie=UTF-8"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3776.0 Safari/537.36")
	client := &http.Client{Timeout: time.Second * 5}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil && data == nil {
		fmt.Println(err)
		return 0, err
	}
	doc, err := htmlquery.Parse(strings.NewReader(string(data)))
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	htmlElement := htmlquery.FindOne(doc, "//*[@id='knowledge-currency__updatable-data-column']/div[1]/div[2]/span[1]")
	response := htmlquery.SelectAttr(htmlElement, "data-value")

	price, err := strconv.ParseFloat(response, 64)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return price, nil
}

func main() {
	price, err := getPriceFromGoogle()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(price)
}
