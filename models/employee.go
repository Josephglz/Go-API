package models

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	UUID	   uint `json:"uuid"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Role       int    `json:"role"`
	Status     int    `json:"status"`
	Phone      string `json:"phone"`
	Address    string `json:"address"`
	City       string `json:"city"`
	State      string `json:"state"`
	ZipCode    string `json:"zipCode"`
	Country    string `json:"country"`
	Department_id int `json:"department"`
}