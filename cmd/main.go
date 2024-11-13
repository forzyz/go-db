package main

import (
	"encoding/json"
	"fmt"

	"github.com/forzyz/go-db/internal/db"
	"github.com/forzyz/go-db/internal/models"
)

func main() {
	dir := "../"

	db, err := db.New(dir, nil)
	if err != nil {
		fmt.Println(err)
	}

	employees := []models.User{
		{
			Name:    "Alice",
			Age:     "29",
			Contact: "1321234",
			Company: "Microsoft",
			Address: models.Address{
				City:    "New York",
				State:   "New York",
				Country: "USA",
				ZipCode: "100001",
			},
		},
		{
			Name:    "Ethan",
			Age:     "22",
			Contact: "3456789",
			Company: "Amazon",
			Address: models.Address{
				City:    "Seattle",
				State:   "Washington",
				Country: "USA",
				ZipCode: "98101",
			},
		},
		{
			Name:    "Mia",
			Age:     "26",
			Contact: "9876543",
			Company: "Apple",
			Address: models.Address{
				City:    "Cupertino",
				State:   "California",
				Country: "USA",
				ZipCode: "95014",
			},
		},
		{
			Name:    "Oliver",
			Age:     "31",
			Contact: "5647382",
			Company: "Samsung",
			Address: models.Address{
				City:    "Seoul",
				State:   "Seoul",
				Country: "South Korea",
				ZipCode: "04524",
			},
		},
		{
			Name:    "Sophia",
			Age:     "24",
			Contact: "8765432",
			Company: "Intel",
			Address: models.Address{
				City:    "Santa Clara",
				State:   "California",
				Country: "USA",
				ZipCode: "95050",
			},
		},
		{
			Name:    "Lucas",
			Age:     "21",
			Contact: "1234567",
			Company: "HP",
			Address: models.Address{
				City:    "Palo Alto",
				State:   "California",
				Country: "USA",
				ZipCode: "94304",
			},
		},
	}

	for _, value := range employees {
		db.Write("users", value.Name, value)
	}

	records, err := db.ReadAll("users")
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println(records)

	allUsers := []models.User{}

	for _, f := range records {
		employeeFound := models.User{}
		if err := json.Unmarshal([]byte(f), &employeeFound); err != nil {
			fmt.Println("Error", err)
		}
		allUsers = append(allUsers, employeeFound)
	}

	fmt.Println(allUsers)

	if err := db.Delete("users", "Alice"); err != nil {
		fmt.Println("Error", err)
	}
}
