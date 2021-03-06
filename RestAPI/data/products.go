package data

import (
	"encoding/json"
	"io"
	"time"
	
	
)

type Product struct {
	ID          int  `json:"id"`
	Name        string `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string   `json:"-"`
	UpdatedOn   string     `json:"-"`
	DeletedOn   string      `json:"-"`
}
func (p *Product) FromJSON(r io.Reader) error {
 e: =json.NewDecoder(r)
 return e.Decode(p)
}
type Products []*Product
func (p  *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func GetProducts() Products {
	return productList
}
func AddProduct(p *Product) {
p.ID =getNextID()
productList = append(productList, p)

}
func getNextID() int {
	lp :=productList[len(productList) - 1]
	return lp.ID + 1
}
var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "GreenTea",
		Description: "good for health",
		Price:       75,
		SKU:         "123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "GreenLand",
		Description: "shows the nature",
		Price:       655,
		SKU:         "xx33",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
