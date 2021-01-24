package gomerchant

type Transaction struct {
	ID        string
	Amount    int
	Currency  string
	Captured  bool
	Paid      bool // if authorized or captured
	Cancelled bool
	Status    string
	CreatedAt *int64
	Params
}
