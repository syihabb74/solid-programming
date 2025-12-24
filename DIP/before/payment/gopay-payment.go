package payment

import "fmt"

type GopayPayment struct {

}

func (gP *GopayPayment) Pay (totalAmount int) {
	fmt.Println("Processing via Gopay with Total = ",totalAmount)
}