package expense

import "time"

type Expense struct {
	Id          int64     `json:"id", omitempty`
	Type        int64     `json:"type"` //expense or income
	Value       float64   `json:"value"`
	Date        time.Time `json:"date", omitempty`
	Comment     string    `json:"comment"`
	Description string    `json:"description"`
}

func NewExpanse() *Expense {
	return &Expense{}

}
