package parse

type Parser interface {
	Parse(filePath string, month int64) []Excrpt
}

type Bank int

const (
	NordeaBank Bank = iota
	SparKronBank
)

func (b Bank) String() string {
	return [...]string{"Nordea", "SparKron"}[b]
}
