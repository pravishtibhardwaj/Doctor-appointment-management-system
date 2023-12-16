package datastore

import (
	"github.com/pravishtibhardwaj/doctor-appointment-management-system/model"

	"gofr.dev/pkg/gofr"
)

type Appointment_i interface {
	// GetByID retrieves a appointment record based on its ID.
	GetByID(ctx *gofr.Context, id string) (*model.Appointment, error)

	GetAllAppointments(ctx *gofr.Context) (*[]model.Appointment,error)
	// Create inserts a new appointment record into the database.
	Create(ctx *gofr.Context, model *model.Appointment) (*model.Appointment, error)
	// Update updates an existing appointment record with the provided information.
	Update(ctx *gofr.Context, model *model.Appointment) (*model.Appointment, error)
	// Delete removes a appointment record from the database based on its ID.
	Delete(ctx *gofr.Context, id int) error
}