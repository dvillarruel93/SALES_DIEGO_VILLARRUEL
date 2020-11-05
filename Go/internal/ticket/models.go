package ticket

import "time"

type Error struct {
	Message    string `json:"message"`
	Error      string `json:"error"`
	StatusCode int    `json:"status_code"`
}

type SaleReceived struct {
	EventID     int64   `json:"event_id"`
	Amount      float64 `json:"amount"`
	Type        string  `json:"sale_type"`
	CountryID   int64   `json:"country_id"`
	CountryName string  `json:"country_name"`
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
	TotalAmount int64  `json:"total_amount"`
}