package utils

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/forzyz/go-db/internal/models"
)

func Stat(path string) (fi os.FileInfo, err error) {
	if fi, err = os.Stat(path); os.IsNotExist(err) {
		fi, err = os.Stat(path + ".json")
	}

	return fi, err
}

func UnmarshalRecords(records []string) []models.User {
	var allUsers []models.User
	for _, record := range records {
		var user models.User
		if err := json.Unmarshal([]byte(record), &user); err != nil {
			fmt.Println("Error unmarshaling record:", err)
		}
		allUsers = append(allUsers, user)
	}
	return allUsers
}
