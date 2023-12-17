# Doctor-appointment-management-system (HTTP-rest-api)
A simple API build to manage the doctor's appointments involving the CRUD operations to create new appointments, get the details of appointments, update details like date and time or status of appointment and delete the appointment once it is being finished.

## Tools and Technologies used :
* This project is built using go language and goFr framework
* Database: MySql database is integrated and data is stored in form of tables
* Testing: Postman is used for testing all api

## Functions :
* ***Create New Appointments:*** This will allow creating new appointments by adding details like the patient's name, doctor's name, Date and time of appointment, and status.Each new appointment will be assigned its own unique id.
* ***Get Details by id:*** This will help in extracting the details of any appointment by inserting the ID of the appointment
* ***List all appointments:*** This will list all the appointments in the database.
* ***Update:*** This will allow us to update the details of any appointment with the given id.
* ***Delete:*** This will allow to delete the existing appointments once they are completed.

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
* To test , use
> go test

## Some Additional Steps
* Create a database using MySQL named "appointments"
  > Create table appointments (id int auto_increment primary key,patient_name varchar(50), doctor_name varchar(50), appointment_time varchar(50), status varchar(50)); 
* Enter the DB details in the .env file in the configs folder
  > HTTP_PORT=9000
  > 
  > DB_HOST=localhost
  > 
  > DB_PORT=3306
  > 
  > DB_USER=your_username
  > 
  > DB_PASSWORD=your_password
  > 
  > DB_NAME=doctor_appointment
  > 
  > DB_DIALECT=mysql

## UML diagrams
* Use Case Diagram:
  
![Untitled (1)](https://github.com/pravishtibhardwaj/Doctor-appointment-management-system/assets/76443753/38626184-f8a6-443f-8312-b3ccc1c05e2d)

***This diagram represents the use case of the application. It includes 2 actors , one is patient and other is admmin. Patient can only view the appointments using the id whereas a admin can create edit and view all the appointments as per request.***

* Sequence Diagram: 

 ![Sequence Diagram Example_ Make Appointment (1)](https://github.com/pravishtibhardwaj/Doctor-appointment-management-system/assets/76443753/79e8ba23-b35f-4341-a35b-9917cc6a0ebd)

 ***Sequence diagram represents the interactions or flow of messages between patient , admin and database.*** 

## Postman testing documentation link:
> https://documenter.getpostman.com/view/21821262/2s9Ykn9N7v

### Some screenshots: 

![Screenshot (145)](https://github.com/pravishtibhardwaj/Doctor-appointment-management-system/assets/76443753/5edf72c5-7583-44f6-9fdb-99f64565c10a)
![Screenshot (146)](https://github.com/pravishtibhardwaj/Doctor-appointment-management-system/assets/76443753/d50bd240-323a-4e26-90c4-a3f13abafd68)







