package service

import (
	"github.com/Carlos6464/imersao22/go-gateway/internal/domain"
	"github.com/Carlos6464/imersao22/go-gateway/internal/dto"
	"github.com/Carlos6464/imersao22/go-gateway/internal/repository"
)

type AccountService struct {
	repository repository.AccountRepository
}

func NewAccountService(repository repository.AccountRepository) *AccountService {
	return &AccountService{repository: repository}
}

func (s *AccountService) CreateAccount(input dto.CreateAccountInput) (*dto.AccountOutput, error) {
	account := dto.ToAccount(input)

	existingAccount, err := s.repository.FindByAPIKey(account.APIKey)

	if err != nil && err != domain.ErrAccountNotFound {
		return nil, err
	}

	if existingAccount != nil {
		return nil, domain.ErrDuplicateAPIKey
	}

	err = s.repository.Save(account)

	if err != nil {
		return nil, err
	}

	output := dto.FromAccount(account)
	return &output, nil

}

func (s *AccountService) UpdateBalance(apikey string, amount float64) (*dto.AccountOutput, error) {
	account, err := s.repository.FindByAPIKey(apikey)
	if err != nil {
		return nil, err
	}

	account.AddBalance(amount)
	err = s.repository.UpdateBalance(account)
	if err != nil {
		return nil, err
	}

	output := dto.FromAccount(account)
	return &output, nil
}

func (s *AccountService) FindByAPIKey(apikey string) (*dto.AccountOutput, error) {
	account, err := s.repository.FindByAPIKey(apikey)
	if err != nil {
		return nil, err
	}
	output := dto.FromAccount(account)
	return &output, nil
}

func (s *AccountService) FindById(id string) (*dto.AccountOutput, error) {
	account, err := s.repository.FindById(id)
	if err != nil {
		return nil, err
	}
	output := dto.FromAccount(account)
	return &output, nil
}
