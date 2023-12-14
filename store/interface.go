package store

import (
	"github.com/pravishtibhardwaj/doctor-appointment-management-system/model"

	"gofr.dev/pkg/gofr"
)

type Student interface {
	
	// Create inserts a new appointment record into the database.
	Create(ctx *gofr.Context, model *model.Appointment) (*model.Appointment, error)

	// GetByID retrieves a appointment detials based on its ID.
	GetByID(ctx *gofr.Context, id string) (*model.Appointment, error)
	
	// Update updates an existing appoinement record with the provided information.
	Update(ctx *gofr.Context, model *model.Appointment) (*model.Appointment, error)
	
	// Delete removes a student record from the database based on its ID.
	Delete(ctx *gofr.Context, id int) error
}
