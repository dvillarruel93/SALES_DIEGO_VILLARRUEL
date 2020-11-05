package ticket

import "time"

type Error struct {
	Error string `json:"error"`
}

type SaleReceived struct {
	EventID     int64   `json:"event_id"`
	Amount      float64 `json:"amount"`
	Type        string  `json:"sale_type"`
	CountryData Country `json:"country_data"`
}

type Country struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	State string `json:"state"`
}

type Sale struct {
	ID          int64     `json:"id"`
	Amount      float64   `json:"amount"`
	SaleType    string    `json:"sale_type"`
	DateCreated time.Time `json:"date_created"`
	EventID     int64     `json:"event_id"`
	CountryID   int64     `json:"country_id"`
	CountryName string    `json:"country_name"`
}

type SalesResume struct {
	CountryName string `json:"country_name"`
	TotalSales  int    `json:"total_sales"`
	TotalAmount int    `json:"total_amoun"`
}