package expense

import "time"

type Expense struct {
	Id          int64     `json:"id"`
	Type        string    `json:"type"` //expense or income
	Value       float64   `json:"value"`
	Date        time.Time `json:"date"`
	Comment     *string   `json:"comment"`
	Description *string   `json:"description"`
}

func NewExpanse() *Expense {
	return &Expense{}

}
