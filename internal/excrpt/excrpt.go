package excrpt

func CreateExcrpt(date, amount, balance float64, description string) Excrpt {
	return Excrpt{
		Date:        date,
		Amount:      amount,
		Description: description,
		Balance:     balance,
	}
}
