package user

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type Service struct {
	userRepository UserRepository
}

func NewService() *Service {
	return &Service{}
}

type InputRegisterUser struct {
	Email    string
	Password string
}

func (s *Service) RegisterUser(ctx context.Context, input *InputRegisterUser) (*User, error) {
	exist := true
	_, err := s.userRepository.GetByEmail(ctx, input.Email)
	if err != nil {
		switch err.(type) {
		case *NotFoundUserRepositoryError:
			exist = false
		default:
			return nil, fmt.Errorf("user service register user err: &v", err)
		}
	}

	if exist {
		return nil, &SameEmailUserAlreadyExistError{fmt.Errorf("user service register user err: same email user already exist")}
	}

	u := &User{
		ID:       uuid.New().String(),
		Email:    input.Email,
		Password: input.Password,
	}

	if err := s.userRepository.Put(ctx, u); err != nil {
		return nil, fmt.Errorf("user service set payment account id err: &v", err)
	}

	return u, nil
}

type InputGetUser struct {
	UserID string
}

func (s *Service) GetUser(ctx context.Context, input *InputGetUser) (*User, error) {
	u, err := s.userRepository.Get(ctx, input.UserID)
	if err != nil {
		switch err.(type) {
		case *NotFoundUserRepositoryError:
			return nil, &NotFoundUserError{err: fmt.Errorf("user service get user err: not found user")}
		default:
			return nil, fmt.Errorf("user service get user err: &v", err)
		}
	}

	return u, nil
}

type InputSetPaymentAccountID struct {
	UserID           string
	PaymentAccountID string
}

func (s *Service) SetPaymentAccountID(ctx context.Context, input *InputSetPaymentAccountID) (*User, error) {
	u, err := s.userRepository.Get(ctx, input.UserID)
	if err != nil {
		switch err.(type) {
		case *NotFoundUserRepositoryError:
			return nil, &NotFoundUserError{err: fmt.Errorf("user service get user err: not found user")}
		default:
			return nil, fmt.Errorf("user service set payment account id err: &v", err)
		}
	}

	u.PaymentAccountID = input.PaymentAccountID

	if err := s.userRepository.Put(ctx, u); err != nil {
		return nil, fmt.Errorf("user service set payment account id err: &v", err)
	}

	return u, nil
}
