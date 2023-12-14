package model

import "time"
type Doctor struct {
	ID      int    `json:"id"`
	D_name  string `json:"name"`
	D_age   int    `json:"age"`
	D_phone string `json:"class"`
}

type Patient struct {
	ID        int    `json:"id"`
	P_name    string `json:"name"`
	P_age     int    `json:"age"`
	P_phone   string `json:"phone"`
	P_address string `json:"address"`
}

type Appointment struct {
	ID   int `json:"id"`
	Patient_name string `json:"pname"`
	Doctor_name string `json:"dname"`
	Apointment_time time.Time `json:"date"`
	Status string `json:"status"`
}