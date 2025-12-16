package order

import "fmt"


type Print struct {

}


func (p *Print) Print (o *Order) {
	p.selectedPrint(o)
}

func (p *Print) printDiscount (totalPrice float64) {
	fmt.Println("==== Discount ====")
	fmt.Println("Price : ", totalPrice)
}

func (p *Print) printNonDiscount (totalPrice float64) {
	fmt.Println("**** NonDiscount ****")
	fmt.Println("Price : ", totalPrice)
}

func (p *Print) selectedPrint (o *Order) {
	if o.HasDiscount{
		p.printDiscount(o.TotalPrice)
	} else {
		p.printNonDiscount(o.TotalPrice)
	}
}