package expense

type Expense struct {
	Type        int64   `json:"type"` //expense or income
	Value       float64 `json:"value"`
	Date        float64 `json:"date"`
	Comment     string  `json:"comment"`
	Description string  `json:"description"`
}

func NewExpanse() *Expense {
	return &Expense{}

}
