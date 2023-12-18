package internal

import "time"

type Report struct {
	ID          int       `json:"id"`
	Income      float32   `json:"income"`
	Expenses    float32   `json:"expenses"`
	Coefficient float32   `json:"coefficient"`
	Date        time.Time `json:"date"`
	OutletId    int       `json:"outlet_id"`
}
