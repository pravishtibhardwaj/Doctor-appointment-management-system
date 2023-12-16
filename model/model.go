package model

type Appointment struct {
	ID               int    `json:"id"`
	Patient_name     string `json:"patient_name"`
	Doctor_name      string `json:"doctor_name"`
	Appointment_time string `json:"appointment_time"`
	Status           string `json:"status"`
}