package model

type Appointment struct {
	ID               int    `json:"id"`
	Patient_name     string `json:"pname"`
	Doctor_name      string `json:"dname"`
	Appointment_time string `json:"time"`
	Status           string `json:"status"`
}