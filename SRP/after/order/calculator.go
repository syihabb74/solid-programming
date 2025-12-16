package order


type Calculator struct {

}


func (c *Calculator) CalculateTotalPrice (o *Order) float64 {

	var sum  float64 = 0;

	for _, v := range o.Products {
		sum += v.Price
	}

	o.TotalPrice = sum

	return sum

}


func (c *Calculator) CalculateDiscount (o *Order) float64 {
	
	totalPrice := c.CalculateTotalPrice(o)
	if totalPrice >= 100000 {
		o.HasDiscount = true
		o.TotalPrice = totalPrice - (0.2 * totalPrice)
		return totalPrice - (0.2 * totalPrice)
	}
	o.HasDiscount = false
	return totalPrice


}