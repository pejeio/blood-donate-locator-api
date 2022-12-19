package configs

import (
	"encoding/csv"
	"os"
)

var AuthUsers map[string]string

func GetAuthUsers() map[string]string {
	return AuthUsers
}

func InitAuthUsers() error {
	file, err := os.Open("auth-users.csv")
	if err != nil {
		return err
	}

	r := csv.NewReader(file)
	r.Comma = ','
	csvUsers, err := r.ReadAll()

	var m = make(map[string]string)
	if err != nil {
		return err
	}
	for _, user := range csvUsers {
		m[user[0]] = user[1]
	}
	AuthUsers = m
	return nil
}
