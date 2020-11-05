package service

import (
	"github.com/dvillarruel93/SALES_DIEGO_VILLARRUEL/internal/ticket"
)

type TicketRepository interface {
	Get() ([]ticket.SalesResume, error)
	Save(saleReceived ticket.Sale) (ticket.Sale, error)
}

type Service struct {
	repository TicketRepository
}

func NewService(ticketRepo TicketRepository) *Service {
	return &Service{
		ticketRepo,
	}
}

func (s Service) Get() ([]ticket.SalesResume, error) {
	value, err := s.repository.Get()

	if err != nil {
		return value, err
	}

	return value, err
}

func (s Service) Save(saleReceived ticket.SaleReceived) (ticket.Sale, error) {
	sale := toSale(saleReceived)
	saleSaved, err := s.repository.Save(sale)

	if err != nil {
		return saleSaved, err
	}

	return saleSaved, nil
}

func toSale(saleReceived ticket.SaleReceived) ticket.Sale {
	return ticket.Sale{
		Amount:      saleReceived.Amount,
		SaleType:    saleReceived.Type,
		EventID:     saleReceived.EventID,
		CountryID:   saleReceived.CountryData.ID,
		CountryName: saleReceived.CountryData.Name,
	}
}