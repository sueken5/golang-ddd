package application

import (
	"context"
	"fmt"

	"github.com/sueken5/golang-ddd/pkg/dddshop/domain/payment"
	"github.com/sueken5/golang-ddd/pkg/dddshop/domain/shop"
	"github.com/sueken5/golang-ddd/pkg/dddshop/domain/user"
)

type DDDShopV1 struct {
	payment *payment.Service
	shop    *shop.Service
	user    *user.Service
}

func NewDDDShopV1() *DDDShopV1 {
	return &DDDShopV1{}
}

type Item shop.Item
type User user.User

func (a *DDDShopV1) ListItems(ctx context.Context) []*Item {
	items := a.ListItems(ctx)
	result := make([]*Item, len(items))
	for i := range result {
		result[i] = (*Item)(items[i])
	}

	return result
}

type InputBuyItem struct {
	ItemID    string
	UserID    string
	RequestID string
}

func (a *DDDShopV1) BuyItem(ctx context.Context, input *InputBuyItem) error {
	//Transaction start もしRequestIDで冪等性の担保をする
	u, err := a.user.GetUser(ctx, &user.InputGetUser{input.UserID})
	if err != nil {
		switch err.(type) {
		case *user.NotFoundUserError:
			return &UnkwonUserError{err: fmt.Errorf("dddshopv1 buy item err: %v", err)}
		default:
			return &UnexpectedError{err: fmt.Errorf("dddshopv1 buy item err: %v", err)}
		}
	}

	if u.PaymentAccountID == "" {
		return &NotPaymentAccountRegisteredError{err: fmt.Errorf("dddshopv1 buy item err: %v", err)}
	}

	item, err := a.shop.BuyItem(ctx, &shop.InputBuyItem{input.ItemID})
	if err != nil {
		switch err.(type) {
		case *shop.ItemIsAlreadySoldError:
			return &ItemIsAlreadySoldError{err: fmt.Errorf("dddshopv1 buy item err: %v", err)}
		default:
			return &UnexpectedError{err: fmt.Errorf("dddshopv1 buy item err: %v", err)}
		}
	}

	if _, err := a.payment.Pay(ctx, &payment.InputPay{AccountID: u.PaymentAccountID, Price: item.Price}); err != nil {
		switch err.(type) {
		case *payment.NotFoundAccountError:
			return &NotPaymentAccountRegisteredError{err: fmt.Errorf("dddshopv1 buy item err: %v", err)}
		default:
			return &UnexpectedError{err: fmt.Errorf("dddshopv1 buy item err: %v", err)}
		}
	}

	//Transaction end
	return nil
}

type InputRegisterUser struct {
	Email    string
	Password string
}

func (a *DDDShopV1) RegisterUser(ctx context.Context, input *InputRegisterUser) (*User, error) {
	u, err := a.user.RegisterUser(ctx, &user.InputRegisterUser{Email: input.Email, Password: input.Password})
	if err != nil {
		switch err.(type) {
		case *user.SameEmailUserAlreadyExistError:
			return nil, &UserAlreadyExistError{fmt.Errorf("dddshopv1 register user err: %v", err)}
		default:
			return nil, &UnexpectedError{err: fmt.Errorf("dddshopv1 register user err: %v", err)}
		}
	}

	return (*User)(u), nil
}

type InputSetupPayment struct {
	UserID           string
	CreditCardNumber string
}

func (a *DDDShopV1) SetupPayment(ctx context.Context, input *InputSetupPayment) (*User, error) {
	//Transaction start
	u, err := a.user.GetUser(ctx, &user.InputGetUser{input.UserID})
	if err != nil {
		switch err.(type) {
		case *user.NotFoundUserError:
			return nil, &UnkwonUserError{err: fmt.Errorf("dddshopv1 setup payment err: %v", err)}
		default:
			return nil, &UnexpectedError{err: fmt.Errorf("dddshopv1 setup payment err: %v", err)}
		}
	}

	account, err := a.payment.RegisterAccount(ctx, &payment.InputRegisterAccount{Email: u.Email, CreditCardNumber: input.CreditCardNumber})
	if err != nil {
		switch err.(type) {
		case *payment.SameEmailAccountAlreadyExistError:
			return nil, &PaymentAlreadySetupError{err: fmt.Errorf("dddshopv1 setup payment err: %v", err)}
		default:
			return nil, &UnexpectedError{err: fmt.Errorf("dddshopv1 setup payment err: %v", err)}
		}
	}

	u, err = a.user.SetPaymentAccountID(ctx, &user.InputSetPaymentAccountID{UserID: u.ID, PaymentAccountID: account.ID})
	if err != nil {
		switch err.(type) {
		case *user.NotFoundUserError:
			return nil, &UnkwonUserError{err: fmt.Errorf("dddshopv1 setup payment err: %v", err)}
		default:
			return nil, &UnexpectedError{err: fmt.Errorf("dddshopv1 setup payment err: %v", err)}
		}
	}
	//Transaction end

	return (*User)(u), nil
}
