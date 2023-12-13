package store

import (
	"encoding/json"
	"io/ioutil"
)

type Product struct {
	ID       string  `json:"id"`
	Category string  `json:"category"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

type Basket struct {
	ID       string    `json:"id"`
	Products []Product `json:"products"`
	Total    float64   `json:"total"`
}

type Customer struct {
	ID        string  `json:"id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Cash      float64 `json:"cash"`
	Basket    Basket  `json:"basket"`
}

func ReadData(filename string) ([]Customer, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var customers []Customer
	err = json.Unmarshal(data, &customers)
	if err != nil {
		return nil, err
	}

	return customers, nil
}
