package data

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"io"
	"time"
)


type Product struct {
	ID 			int 			`json:"ID"`
	Name 		string			`json:"Name"         validate:"required"`
	Description string			`json:"Description"`
	Price 	  	float32			`json:"Price"        validate:"required"`
	SKU		  	string			`json:"SKU"          validate:"required"`
	CreatedOn 	string			`json:"-"`
	UpdatedON 	string			`json:"-"`
	DeletedOn 	string			`json:"-"`
}

func (p *Product) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}

func (p *Product) Validate() error {
	validate := validator.New()
	return validate.Struct(p)
}

// Products is a collection of Product
type Products []*Product

// ToJSON serializes the contents of the collection to JSON
// NewEncoder provides better performance than json.Unmarshal as it does not
// have to buffer the output into an in memory slice of bytes
// this reduces allocations and the overheads of the service
// https://golang.org/pkg/encoding/json/#NewEncoder
func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func GetProducts() Products {
	return productList
}

func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

func UpdateProduct(id int, p*Product) error {
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}
	p.ID = id
	productList[pos] = p
	return nil
}

var ErrProductNotFound = fmt.Errorf("Product not found")

func findProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}

	return nil, -1, ErrProductNotFound
}


func getNextID() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
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
		Name:        "Espresso",
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
