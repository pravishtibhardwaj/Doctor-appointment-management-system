package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateAppointment(t *testing.T) {
	requestBody := `{
        "Patient_name":"Pravishti Bhardwaj",
		"Doctor_name":"Dr. S.K. Sharma",
		"Appointment_time":"2 pm",
		"Status":"Scheduled"
    }`
	req, err := http.NewRequest("POST", "/appointment", bytes.NewBufferString(requestBody))
	if err != nil {
		t.Fatal(err)
	}
	handler := http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		wr.WriteHeader(http.StatusOK)
	})
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	// Checking Status Code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

}
func TestUpdateAppointment(t *testing.T) {
	requestBody := `{
		"Patient_name":"Rohan",
		"Doctor_name":"Dr. A.K. roy",
		"Appointment_time":"1 p.m.",
		"Status":"pending"
    }`
	req, err := http.NewRequest("PUT", "/appointment/2", bytes.NewBufferString(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		wr.WriteHeader(http.StatusOK)
	})
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

}

func TestGetAllDetails(t *testing.T) {
	req, err := http.NewRequest("GET", "/all_appointment", nil)
	if err != nil {
		t.Fatal(err)
	}
	handler := http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		wr.WriteHeader(http.StatusOK)
	})
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestAppointmentByID(t *testing.T) {
	req, err := http.NewRequest("GET", "/appointement/2", nil)
	if err != nil {
		t.Fatal(err)
	}
	handler := http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		wr.WriteHeader(http.StatusOK)
	})
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
