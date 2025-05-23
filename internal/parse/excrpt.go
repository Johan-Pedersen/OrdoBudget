package parse

type Excrpt struct {
	Date        string
	Amount      float64
	Description string
	Balance     float64
}

func CreateExcrpt(amount, balance float64, date, description string) Excrpt {
	return Excrpt{
		Date:        date,
		Amount:      amount,
		Description: description,
		Balance:     balance,
	}
}

func (e Excrpt) Equals(exc Excrpt) bool {
	return exc.Date == e.Date &&
		e.Amount == exc.Amount &&
		e.Balance == exc.Balance &&
		e.Description == exc.Description
}
