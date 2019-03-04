package main

type Employee struct {
	Id			string	`form:"id" json:"id"`
	FullName	string	`form:"fullname" json:"fullname"`
	Position	string	`form:"position" json:"position"`
	EMPCode		string	`form:"empcode" json:"empcode"`
	Mobile		string	`form:"mobile" json:"mobile"`
}

type Response struct {
	Status	int		`json:"status"`
	Message	string	`json:"message"`
	Data	[]Employee
}