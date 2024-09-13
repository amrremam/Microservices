package data

import (
	"encoding/json"
	"io"
	"time"
)

type Product struct {
	ID int `json:"ID"`
	Name string	`json:"Name"`
	Description string	`json:"Description"`
	Price float32	`json:"Price"`
	SKU string	`json:"SKU"`
	CreatedOn string	`json:"-"`
	UpdatedON string	`json:"-"`
	DeletedOn string	`json:"-"`
}


type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func GetProducts() Products {
	return productList
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
		CreatedOn: time.Now().UTC().String(),
		UpdatedON: time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Esspresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd34",
		CreatedOn: time.Now().UTC().String(),
		UpdatedON: time.Now().UTC().String(),
	},
	&Product{
		ID:          3,
		Name:        "Eroresso",
		Description: "Long",
		Price:       2.22,
		SKU:         "fjd322",
		CreatedOn: time.Now().UTC().String(),
		UpdatedON: time.Now().UTC().String(),
	},
}
