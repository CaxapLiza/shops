package internal

type Outlet struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	Address      string  `json:"address"`
	PlanedProfit float32 `json:"planed_profit"`
	OwnerId      int     `json:"owner_id"`
}
