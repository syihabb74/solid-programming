package payment

type PaymentProcess interface {
	Pay(amount int)
}

type Payment struct {
	TotalAmount int
}

func (p *Payment) PaymentProcess (pP PaymentProcess) {
	pP.Pay(p.TotalAmount)
}
