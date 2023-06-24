package service

import (
	"time"

	"github.com/vietbui1502/RestAPIGolang/domain"
	"github.com/vietbui1502/RestAPIGolang/dto"
)

type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, error)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (s DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, error) {
	a := domain.Account{
		AccountId:   "",
		CustomerId:  req.CustomerId,
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
	}

	newAccount, err := s.repo.Save(a)
	if err != nil {
		return nil, err
	}

	respDto := newAccount.ToAccountResponseDTO()
	return &respDto, nil
}

func NewAccountService(repository domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo: repository}
}
