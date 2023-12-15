package datastore

import (
	"database/sql"

	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"

	"github.com/pravishtibhardwaj/doctor-appointment-management-system/model"
)

type appoint struct{}

func New() *appoint {
	return &appoint{}
}

func (s *appoint) GetByID(ctx *gofr.Context, id string) (*model.Appointment, error) {
	var resp model.Appointment

	err := ctx.DB().QueryRowContext(ctx, " SELECT id,pname,dname,time,status FROM appointments where id=?", id).
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

func (s *appoint) Create(ctx *gofr.Context, appoint *model.Appointment) (*model.Appointment, error) {
	

	var resp model.Appointment

    result, err := ctx.DB().ExecContext(ctx, "INSERT INTO appointments (pname,dname,time,status) VALUES (?,?,?,?)",appoint.Patient_name, appoint.Doctor_name, appoint.Appointment_time,appoint.Status)
    if err != nil {
        return &model.Appointment{}, errors.DB{Err: err}
    }

    lastInsertID, err := result.LastInsertId()
	
    if err != nil {
        return &model.Appointment{}, errors.DB{Err: err}
    }
// ctx.Logger.Log(lastInsertID)
    err = ctx.DB().QueryRowContext(ctx, "SELECT id, pname, dname, time,status FROM appointments WHERE id = ?", lastInsertID).Scan(&resp.ID, &resp.Patient_name, &resp.Doctor_name,&resp.Appointment_time, &resp.Status)
    if err != nil {
        return &model.Appointment{}, errors.DB{Err: err}
    }

    return &resp, nil
}

func (s *appoint) Update(ctx *gofr.Context, appoint *model.Appointment) (*model.Appointment, error) {
	_, err := ctx.DB().ExecContext(ctx, "UPDATE appointments SET pname=?,dname=?,time=?,status=? WHERE id=?",
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