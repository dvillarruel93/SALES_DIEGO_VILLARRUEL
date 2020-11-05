package api

import (
	"github.com/dvillarruel93/SALES_DIEGO_VILLARRUEL/internal/http/handler/middleware"
	"github.com/dvillarruel93/SALES_DIEGO_VILLARRUEL/internal/http/handler/ticket"
	"github.com/dvillarruel93/SALES_DIEGO_VILLARRUEL/internal/platform/database"
	"github.com/dvillarruel93/SALES_DIEGO_VILLARRUEL/internal/ticket/repository"
	"github.com/dvillarruel93/SALES_DIEGO_VILLARRUEL/internal/ticket/service"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func Start() {
	Router = gin.Default()

	load()

	Router.Run(":8080")
}

func load() {
	// Load database
	db := database.NewDatabase()

	// Load repository
	ticketRepo := repository.NewTicketRepository(db)

	// Load service with repositories
	ticketService := service.NewService(ticketRepo)

	// Load handler with services
	ticketHandler := ticket.NewTicketHandler(ticketService)

	Router.GET("/api/v1/stats", middleware.AuthMiddleware(), ticketHandler.GetTicket)
	Router.POST("/api/v1/sales", middleware.AuthMiddleware(), ticketHandler.SaveTicket)
}