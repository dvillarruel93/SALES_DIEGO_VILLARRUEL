package repository

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/dvillarruel93/SALES_DIEGO_VILLARRUEL/internal/platform/database"
	"github.com/dvillarruel93/SALES_DIEGO_VILLARRUEL/internal/ticket"
)

const DBFormatDateTime = "2006-01-02 15:04:05"

type TicketRepository struct {
	database database.Database
}

func NewTicketRepository(db database.Database) TicketRepository {
	return TicketRepository{
		database: db,
	}
}

func (r TicketRepository) Get() ([]ticket.SalesResume, error) {
	query := "SELECT country_name, COUNT(*) AS total_sales, SUM(amount) AS total_amount FROM sale GROUP BY country_name ORDER BY total_sales DESC"
	rows, err := r.database.SelectMultiple(query, nil)

	if err != nil {
		log.Printf("error doing multiple select: %s", err.Error())
		return []ticket.SalesResume{}, errors.New("error doing multiple select")
	}

	defer rows.Close()
	var salesResume []ticket.SalesResume

	for rows.Next() {
		var saleResume ticket.SalesResume
		if err := rows.Scan(&saleResume.CountryName, &saleResume.TotalSales, &saleResume.TotalAmount); err != nil {
			log.Printf("error scanning row: %s", err.Error())
			return salesResume, errors.New("error scanning row")
		}

		salesResume = append(salesResume, saleResume)
	}

	return salesResume, nil
}

func (r TicketRepository) Save(sale ticket.Sale) (ticket.Sale, error) {
	now := time.Now()
	nowFormatted := now.Format(DBFormatDateTime)
	query := fmt.Sprintf("INSERT INTO sale (amount, sale_type, date_created, event_id, country_id, country_name) VALUES (%f, '%s', '%s', %d, %d, '%s')",
		sale.Amount,
		sale.SaleType,
		nowFormatted,
		sale.EventID,
		sale.CountryID,
		sale.CountryName)
	sqlResult, err := r.database.ExecuteQuery(query)
	if err != nil {
		log.Printf("error executing query: %s", err.Error())
		return sale, errors.New("error executing query")
	}

	sale.DateCreated = now

	lastInsertedID, err := sqlResult.LastInsertId()
	if err == nil {
		sale.ID = lastInsertedID
	}

	return sale, nil
}