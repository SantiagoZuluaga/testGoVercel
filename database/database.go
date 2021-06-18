package database

import (
	"fmt"
)

type Currency struct {
	Symbol  string `json:"symbol"`
	Price   int    `json:"price"`
	Updated int    `json:"updated"`
}

func GetCurrencyPrice() (Currency, error) {

	var currency Currency
	db, err := GetConnection()
	if err != nil {
		fmt.Println("ERROR GET DATABASE", err)
		return Currency{}, err
	}

	row := db.QueryRow("SELECT symbol, price, updated FROM currency")
	err = row.Scan(
		&currency.Symbol,
		&currency.Price,
		&currency.Updated,
	)
	if err != nil {
		return Currency{}, err
	}

	return currency, nil
}

func InsertCurrency(price int, time int) (Currency, error) {
	statement, _ := db.Prepare("INSERT INTO currency (symbol, price, updated) VALUES (?, ?, ?)")
	if err != nil {
		fmt.Println("ERROR INSERT CURRENCY", err)
		return Currency{}, err
	}
	statement.Exec("COP", price, time)
	return Currency{
		Symbol:  "COP",
		Price:   price,
		Updated: time,
	}, nil
}

func UpdateCurrency(price int, time int) (Currency, error) {
	statement, _ := db.Prepare("UPDATE currency set price = ?, updated = ? where symbol = 'COP'")
	response, _ := statement.Exec(price, time)
	affect, _ := response.RowsAffected()
	fmt.Println(affect)

	return Currency{
		Symbol:  "COP",
		Price:   price,
		Updated: time,
	}, nil
}
