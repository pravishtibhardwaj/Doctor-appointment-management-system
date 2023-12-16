package datastore

import (
	"database/sql"

	"github.com/pravishtibhardwaj/doctor-appointment-management-system/model"
	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
)

type appoint struct{}

func New() *appoint {
	return &appoint{}
}

func (s *appoint) GetByID(ctx *gofr.Context, id string) (*model.Appointment, error) {
	var resp model.Appointment

	  err := ctx.DB().QueryRowContext(ctx, "SELECT id,patient_name, doctor_name, appointment_time,status FROM appointments where id=?", id).
		Scan(&resp.ID, &resp.Patient_name, &resp.Doctor_name, &resp.Appointment_time,&resp.Status)
	switch err {
	case sql.ErrNoRows:
		return &model.Appointment{}, errors.EntityNotFound{Entity: "appointments", ID: id}
	case nil:
		return &resp, nil
	default:
		return &model.Appointment{}, err
	}
}
func (s *appoint) GetAllAppointments(ctx *gofr.Context) (*[]model.Appointment, error) {
	var appointments []model.Appointment

	rows, err := ctx.DB().QueryContext(ctx, "SELECT id, patient_name, doctor_name, appointment_time, status FROM appointments")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var appointment model.Appointment
		err := rows.Scan(&appointment.ID, &appointment.Patient_name, &appointment.Doctor_name, &appointment.Appointment_time, &appointment.Status)
		if err != nil {
			return nil, err
		}
		appointments = append(appointments, appointment)
	}

	return &appointments, nil
}


func (s *appoint) Create(ctx *gofr.Context, appoint *model.Appointment) (*model.Appointment, error) {
	

	var resp model.Appointment

    result, err := ctx.DB().ExecContext(ctx, "INSERT INTO appointments (patient_name, doctor_name, appointment_time,status) VALUES (?,?,?,?)",appoint.Patient_name, appoint.Doctor_name, appoint.Appointment_time,appoint.Status)
    if err != nil {
        return &model.Appointment{}, errors.DB{Err: err}
    }

    lastInsertID, err := result.LastInsertId()
	
    if err != nil {
        return &model.Appointment{}, errors.DB{Err: err}
    }
// ctx.Logger.Log(lastInsertID)
    err = ctx.DB().QueryRowContext(ctx, "SELECT id, patient_name, doctor_name, appointment_time,status FROM appointments WHERE id = ?", lastInsertID).Scan(&resp.ID, &resp.Patient_name, &resp.Doctor_name,&resp.Appointment_time, &resp.Status)
    if err != nil {
        return &model.Appointment{}, errors.DB{Err: err}
    }

    return &resp, nil
}

func (s *appoint) Update(ctx *gofr.Context, appoint *model.Appointment) (*model.Appointment, error) {
	_, err := ctx.DB().ExecContext(ctx, "UPDATE appointments SET patient_name=?, doctor_name=?, appointment_time=?,status=? WHERE id=?",
		appoint.Patient_name, appoint.Doctor_name, appoint.Appointment_time,appoint.Status, appoint.ID)
	if err != nil {
		return &model.Appointment{}, errors.DB{Err: err}
	}

	return appoint, nil
}

func (s *appoint) Delete(ctx *gofr.Context, id int) error {
	_, err := ctx.DB().ExecContext(ctx, "DELETE FROM appointments where id=?", id)
	if err != nil {
		return errors.DB{Err: err}
	}

	return nil
}