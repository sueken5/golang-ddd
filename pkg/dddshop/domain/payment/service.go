package payment

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Service struct {
	accountRepository AccountRepository
	receiptRepository ReceiptRepository
}

func NewService() *Service {
	return &Service{}
}

type InputRegisterAccount struct {
	Email            string
	CreditCardNumber string
}

func (s *Service) RegisterAccount(ctx context.Context, input *InputRegisterAccount) (*Account, error) {
	exist := true
	_, err := s.accountRepository.GetByEmail(ctx, input.Email)
	if err != nil {
		switch err.(type) {
		case *NotFoundAccountRepositoryError:
			exist = false
		default:
			return nil, fmt.Errorf("payment service register account err: %v", err)
		}
	}

	if exist {
		return nil, &SameEmailAccountAlreadyExistError{fmt.Errorf("payment service register account err: same email account is already exist")}
	}

	a := &Account{
		ID:               uuid.New().String(),
		Email:            input.Email,
		CreditCardNumber: input.CreditCardNumber,
	}

	if err := s.accountRepository.Put(ctx, a); err != nil {
		return nil, fmt.Errorf("payment service register account err: %v", err)
	}

	return a, nil
}

type InputPay struct {
	AccountID string
	RequestID string
	Price     int
}

func (s *Service) Pay(ctx context.Context, input *InputPay) (*Receipt, error) {
	_, err := s.accountRepository.Get(ctx, input.AccountID)
	if err != nil {
		switch err.(type) {
		case *NotFoundAccountRepositoryError:
			return nil, &NotFoundAccountError{fmt.Errorf("payment service pay err: not found account")}
		default:
			return nil, fmt.Errorf("payment service pay err: %v", err)
		}
	}

	r := &Receipt{
		ID:        input.RequestID,
		AccountID: input.AccountID,
		Price:     input.Price,
		CreatedAt: time.Now(),
	}

	if err := s.receiptRepository.Put(ctx, r); err != nil {
		return nil, fmt.Errorf("payment service pay err: %v", err)
	}

	return r, nil
}
