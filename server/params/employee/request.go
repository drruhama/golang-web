package params

import (
	"go-web-template/server/models"
	"time"
)

// berfungsi untuk handle request
// untuk membuat employee baru
type EmployeeCreate struct {
	NIP     string
	Name    string
	Address string
}

// method ini berfungsi untuk mengubah data params ke models
func (e *EmployeeCreate) ParseToModel() *models.Employee {
	employee := models.NewEmployee()
	employee.Address = e.Address
	employee.Nama = e.Name
	employee.NIP = e.NIP
	return employee
}

// ini berfungsi untuk handle request
// jika ada update employee
type EmployeeUpdate struct {
	ID        int
	NIP       string
	Name      string
	Address   string
	UpdatedAt time.Time
}

// method ini berfungsi untuk mengubah data params ke models
func (e *EmployeeUpdate) ParseToModel() *models.Employee {
	return &models.Employee{
		ID:        e.ID,
		NIP:       e.NIP,
		Nama:      e.Name,
		Address:   e.Address,
		UpdatedAt: time.Now(),
	}
}
