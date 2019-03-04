package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func getEmployees(w http.ResponseWriter, r *http.Request) {
	var employee Employee
	var employees []Employee
	var response Response

	db := connect()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM employees")
		if err != nil {
			log.Print(err)
		}

		for rows.Next() {
			if err := rows.Scan(&employee.Id, &employee.FullName, &employee.Position, &employee.EMPCode, &employee.Mobile); err != nil {
				log.Fatal(err.Error())
			} else {
				employees = append(employees, employee)
			}
		}

		response.Status = 1
		response.Message = "Success"
		response.Data = employees

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
}

func postEmployee(w http.ResponseWriter, r *http.Request) {
	var response Response

		db := connect();
		defer db.Close()

		err := r.ParseMultipartForm(4096)

			if err != nil {
				panic(err)
			}

		fullname := r.FormValue("fullname")
		position := r.FormValue("position")
		empcode  := r.FormValue("empcode")
		mobile	 := r.FormValue("mobile")

		_, err = db.Exec("INSERT INTO employees VALUES (null, ?, ?, ?, ?)", fullname, position, empcode, mobile)
			
			if err != nil {
				log.Print(err)
			}
	response.Status = 1
	response.Message = "success"
	log.Print("Inserting data to database..")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}