package models

import "encoding/json"

type Address struct {
	City    string
	State   string
	Country string
	ZipCode json.Number
}

type User struct {
	Name    string
	Age     json.Number
	Contact string
	Company string
	Address Address
}
