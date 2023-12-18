package internal

type Employee struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Itn       string `json:"itn"`
	Passport  string `json:"passport"`
	Snils     string `json:"snils"`
	Phone     string `json:"phone"`
	AccountId int    `json:"account_id"`
}
