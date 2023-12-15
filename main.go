package main

import (
	"github.com/pravishtibhardwaj/doctor-appointment-management-system/datastore"

	"github.com/pravishtibhardwaj/doctor-appointment-management-system/handler"

	"gofr.dev/pkg/gofr"
)

func main() {
	app := gofr.New()

	s := datastore.New()
	h := handler.New(s)

	app.GET("/appointment/{id}", h.GetByID)
	app.POST("/appointment", h.Create)
	app.PUT("/appointment/{id}", h.Update)
	app.DELETE("/appointment/{id}", h.Delete)

	// starting the server on a custom port
	app.Server.HTTP.Port = 9000
	app.Start()
}
