package main

import (
	"DIP/after/payment"
	_ "DIP/before/payment"
)

func main () {
	// gopay := &payment.GopayPayment{}
	// paymentMachine := &payment.Payment{TotalAmount: 50000};
	// paymentMachine.PaymentProcess(gopay) // this is violate dependency inversion principle causes 
	// the payment machine a.k.a Payment have an konkrete implementation now imagine we want to change
	// the payment machine with DanaPayment we have to go to source file payment and change with explicit 
	// way again and assume 2 days after we want to change again with OVO payment we have to change again
	// and go to source file that need to change 


	gopay := &payment.GopayPayment{}
	dana := &payment.DanaPayment{}
	paymentMachine := &payment.Payment{TotalAmount: 100000}
	paymentMachine.PaymentProcess(gopay)
	paymentMachine.PaymentProcess(dana) // this is the solution and called dependency inversion principle

}