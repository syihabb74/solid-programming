package before

import "fmt"

type Order struct {
	ID string
	Products []*Product
	Shipping string
}

func (o *Order) CalculateTotalPriceAndDiscount () float64 {
	var sum float64 = 0;
	for _, v := range o.Products {
		sum += v.Price
	}
	if sum >= 100000 {
		return sum - (0.2 * sum)
	}
	return sum
} // violate srp because it's different responsibility with what order supposed to be
// now suppose we want to save original price for different purpose so we can't do this
// because it has 2 different reason to change and will affect to original price we want to save before
// because the price will be unconsistent if the sum above 100k 

func (p *Product) PrintProduct () {
	fmt.Println("=== ", p.Name, " ===")
	fmt.Println("Price : ", p.Price)
	fmt.Println("ID : ", p.ID)
}

func (o *Order) PrintAll () {
	for _ , v := range o.Products {
		v.PrintProduct()
	}
} // violate srp because it's different responsibility with what order supposed to be

func (o *Order) Save (p *Product) {
	o.Products = append(o.Products, p)
}


type Product struct {
	ID string
	Name string
	Price float64
}

