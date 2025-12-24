package payment


type Payment struct {
	TotalAmount int
}

func (p *Payment) PaymentProcess (gP *GopayPayment) {
	gP.Pay(p.TotalAmount)
}