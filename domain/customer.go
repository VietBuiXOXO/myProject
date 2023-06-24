package domain

import "github.com/vietbui1502/RestAPIGolang/dto"

type Customer struct {
	Id          string `db:"customer_id"`
	Name        string
	City        string
	ZipCode     string
	DateofBirth string `db:"date_of_birth"`
	Status      string
}

func (c Customer) ToDTO() dto.CustomerResponse {
	statusAsText := "active"
	if c.Status == "0" {
		statusAsText = "inactive"
	}

	response := dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		ZipCode:     c.ZipCode,
		DateofBirth: c.DateofBirth,
		Status:      statusAsText,
	}

	return response
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
	FindCustomerbyID(string) (*Customer, error)
}
