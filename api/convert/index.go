package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	htmlquery "github.com/antchfx/htmlquery"
)

func getPriceFromGoogle() (int, error) {
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

	price, err := strconv.Atoi(response)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return price, nil
}

func Handler(w http.ResponseWriter, r *http.Request) {

	/*
		now := time.Now()
		unixNano := now.UnixNano()
		umillisec := unixNano / 1000000

		 currency, err := database.GetCurrencyPrice()
		 if err != nil {
			 price, err := getPriceFromGoogle()
			 if err != nil {
				 fmt.Println(err)
			 }
			 newcurrency, err := database.InsertCurrency(price, int(umillisec))
			 if err != nil {
				 fmt.Println(err)
			 }
			 w.Header().Set("Content-Type", "application/json")
			 json.NewEncoder(w).Encode(newcurrency)
			 return
		 }

		 if currency.Updated > int(umillisec)-43200*1000 {
			 w.Header().Set("Content-Type", "application/json")
			 json.NewEncoder(w).Encode(currency)
			 return
		 }



		 if currency.Price != price {
			 newcurrency, err := database.UpdateCurrency(price, int(umillisec))
			 if err != nil {
				 fmt.Println(err)
			 }
			 w.Header().Set("Content-Type", "application/json")
			 json.NewEncoder(w).Encode(newcurrency)
			 return
		 }
	*/

	price, err := getPriceFromGoogle()
	if err != nil {
		fmt.Println(err)
	}

	parm := r.URL.Query()["amount"][0]
	amount, err := strconv.Atoi(parm)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(price * amount)
}
