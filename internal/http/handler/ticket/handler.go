package ticket

import (
	"errors"
	"fmt"
	"github.com/dvillarruel93/SALES_DIEGO_VILLARRUEL/internal/ticket"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type TicketService interface {
	Get() ([]ticket.SalesResume, error)
	Save(saleReceived ticket.SaleReceived) (ticket.Sale, error)
}

type TicketHandler struct {
	TicketService
}

func NewTicketHandler (service TicketService) *TicketHandler {
	return &TicketHandler{
		service,
	}
}

func (h TicketHandler) GetTicket(c *gin.Context) {
	ticketResponse, err := h.Get()

	if err != nil {
		c.JSON(http.StatusInternalServerError, ticket.Error{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, ticketResponse)
}

func (h TicketHandler) SaveTicket(c *gin.Context) {
	request := ticket.SaleReceived{}

	if err := c.BindJSON(&request); err != nil {
		errMsg := fmt.Sprintf("error binding body: %s", err.Error())
		log.Print(errMsg)
		c.JSON(http.StatusInternalServerError, err)

		return
	}

	err := validateRequest(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, ticket.Error{Error: err.Error()})
		return
	}

	ticketResponse, err := h.Save(request)

	if err != nil {
		c.JSON(http.StatusInternalServerError, ticket.Error{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, ticketResponse)
}

func validateRequest(request ticket.SaleReceived) error {
	if request.Amount == 0 {
		return errors.New("amount is a mandatory param")
	}

	if request.Type == "" {
		return errors.New("sale_type is a mandatory param")
	}

	if request.EventID == 0 {
		return errors.New("event_id is a mandatory param")
	}

	if request.CountryData.ID == 0 {
		return errors.New("country.id is a mandatory param")
	}

	if request.CountryData.Name == "" {
		return errors.New("country.name is a mandatory param")
	}

	return nil
}