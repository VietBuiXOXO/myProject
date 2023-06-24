package service

import (
	"github.com/vietbui1502/RestAPIGolang/domain"
	"github.com/vietbui1502/RestAPIGolang/dto"
	"github.com/vietbui1502/RestAPIGolang/logger"
)

type CustomerService interface {
	GetAllCustomer() ([]dto.CustomerResponse, error)
	GetCustomer(string) (*dto.CustomerResponse, error)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer() ([]dto.CustomerResponse, error) {
	customers, err := s.repo.FindAll()

	if err != nil {
		logger.Error("GetAllCustomer error")
		return nil, err
	}

	customersResponse := make([]dto.CustomerResponse, 0)

	for i := 1; i < len(customers); i++ {
		c := customers[i].ToDTO()
		customersResponse = append(customersResponse, c)
	}

	return customersResponse, nil
}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, error) {
	c, err := s.repo.FindCustomerbyID(id)
	if err != nil {
		logger.Error("customer service get customer by id error")
		return nil, err
	}

	response := c.ToDTO()

	return &response, nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}
}
