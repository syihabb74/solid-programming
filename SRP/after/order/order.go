package order

import (
	"SRP/after/product"
	"slices"
)


type Order struct {
	ID string
	Products []*product.Product
	Shipping string
	TotalPrice float64
	HasDiscount bool
}


func (o *Order) Save(p *product.Product) {
	o.Products = append(o.Products, p)
}

func (o *Order) targetDelete (id string) int { // helpers function
	for i, v := range o.Products {
		if v.ID == id {
			return i
		}
	}
	return -1
} 

func (o *Order) Delete(id string) {
	target := o.targetDelete(id);
	o.Products = slices.Delete(o.Products, target, target + 1)
}



