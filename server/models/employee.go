package models

import (
	"time"
)

type Employee struct {
	ID        int
	NIP       string
	Nama      string
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewEmployee() *Employee {
	return &Employee{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
