package payment

import "fmt"

type DanaPayment struct {

}

func (dP *DanaPayment) Pay (totalAmount int) {
	fmt.Println("Processing via Gopay with Total = ",totalAmount)
}