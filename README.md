# Doctor-appointment-management-system
A simple API build to manage the doctor's appointments involves the CRUD operations to create new appointments, get the details of appointments, update details like date and time or status of appointment and delete the appointment once it is being finished.

## Tools and Technologies used :
* This project is built using go language and goFr framework
* Database: MySql database is integrated and data is stored in form of tables
* Testing: Postman is used for testing all api

## Functions :
* Create New Appointments: This will allow creating new appointments by adding details like the patient's name, doctor's name, Date and time of appointment, and status.Each new appointment will be assigned its own unique id.
* Get Details: This will help in extracting the details of any appointment by inserting the ID of the appointment
* Update : This will allow us to update details of any appointment with given id.
* Delete : This will allow to delete the existing appointments once they are completed.

## Steps to run the project
* install the latest version of go onto your system
* Clone this repository using:
> git clone https://github.com/pravishtibhardwaj/Doctor-appointment-management-system.git
* Run the following command to add goFr package:
> go get gofr.dev
* To download and sync the required modules run the following command:
> go mod tidy
* To run the server, use the command:
> go run main.go

## Some Additional Steps
* Create a database using mysql named "appointments"
* Enter the DB details in the .env file in the configs folder

## UML diagrams
* Use Case Diagram:
  
![Untitled (1)](https://github.com/pravishtibhardwaj/Doctor-appointment-management-system/assets/76443753/38626184-f8a6-443f-8312-b3ccc1c05e2d)





