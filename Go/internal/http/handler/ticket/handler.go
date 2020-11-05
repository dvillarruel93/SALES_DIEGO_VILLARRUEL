package ticket

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/dvillarruel93/SALES_DIEGO_VILLARRUEL/internal/ticket"
	"github.com/gin-gonic/gin"
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
		c.JSON(http.StatusInternalServerError, ticket.Error{Message: err.Error(), Error: "internal server error", StatusCode: http.StatusInternalServerError})
		return
	}

	if len(ticketResponse) == 0 {
		c.JSON(http.StatusNotFound, ticket.Error{Message: "no sales", Error: "not found", StatusCode: http.StatusNotFound})
		return
	}

	c.JSON(http.StatusOK, ticketResponse)
}

func (h TicketHandler) SaveTicket(c *gin.Context) {
	request := ticket.SaleReceived{}

	err := c.BindJSON(&request)
	if err != nil {
		errMsg := fmt.Sprintf("error binding request: %s", err.Error())

		c.JSON(http.StatusInternalServerError, ticket.Error{Message: errMsg, Error: "internal server error", StatusCode: http.StatusInternalServerError})

		return
	}

	err = validateRequest(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, ticket.Error{Message: err.Error(), Error: "bad request", StatusCode: http.StatusBadRequest})
		return
	}

	ticketResponse, err := h.Save(request)

	if err != nil {
		c.JSON(http.StatusInternalServerError, ticket.Error{Message: err.Error(), Error: "internal server error", StatusCode: http.StatusInternalServerError})
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

	if request.CountryID == 0 {
		return errors.New("country_id is a mandatory param")
	}

	if request.CountryName == "" {
		return errors.New("country_name is a mandatory param")
	}

	return nil
}