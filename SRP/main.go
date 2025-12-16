package main

import (
	"SRP/after/order"
	"SRP/after/product"
	_ "SRP/before"
	"fmt"
)

func main () {
	// before SRP
	// order := before.Order{};
	// p1 := before.Product{
	// 	ID: "X100",
	// 	Name: "Sunlight",
	// 	Price: 1500,
	// }
	// p2 := before.Product{
	// 	ID: "X101",
	// 	Name: "Bimoli",
	// 	Price: 3500,
	// }
	// p3 := before.Product{
	// 	ID: "X102",
	// 	Name: "Sabun Cuci Muka",
	// 	Price: 5500,
	// }

	// order.Save(&p1)
	// order.Save(&p2)
	// order.Save(&p3)

	// fmt.Println(order)
	// order.CalculateTotalPriceAndDiscount()


	// after SRP
	order1 := &order.Order{}
	p1 := &product.Product{
		ID: "X100",
		Name: "Sunlight",
		Price: 150000,
	}
	p2 := &product.Product{
		ID: "X101",
		Name: "Bimoli",
		Price: 3500,
	}
	p3 := &product.Product{
		ID: "X102",
		Name: "Sabun Cuci Muka",
		Price: 5500,
	}
	order1.Save(p1)
	order1.Save(p2)
	order1.Save(p3)
	fmt.Println(order1)


	calculator := &order.Calculator{}
	fmt.Println(calculator.CalculateTotalPrice(order1)) // if we want to save data for different purpose use here first
	fmt.Println(calculator.CalculateDiscount(order1)) // and here the final price to the customer
	print := &order.Print{};
	print.Print(order1)




	

}