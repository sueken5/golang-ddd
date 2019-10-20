package http

import (
	"context"
	"fmt"
	"net/http"

	"github.com/sueken5/golang-ddd/pkg/dddshop/application"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

type Server struct {
	usecase Usecase
	server  *http.Server
}

type Usecase interface {
	RegisterUser(ctx context.Context, input *application.InputRegisterUser) (*application.User, error)
	SetupPayment(ctx context.Context, input *application.InputSetupPayment) (*application.User, error)
	ListItems(ctx context.Context) ([]*application.Item, error)
	BuyItem(ctx context.Context, input *application.InputBuyItem) error
}

func NewServer() *Server {
	server := &Server{}
	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", 8080),
		Handler: server.router(),
	}

	server.server = s

	return server
}

func (s *Server) router() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/users", s.registerUser()).Methods(http.MethodPost)
	router.HandleFunc("/users/payment", s.setupPayment()).Methods(http.MethodPut)
	router.HandleFunc("/items", s.listItems()).Methods(http.MethodGet)
	router.HandleFunc("/items/buy", s.buyItem()).Methods(http.MethodPost)

	n := negroni.New()
	n.UseHandler(router)

	return n
}

func (s *Server) Run() error {
	if err := s.server.ListenAndServe(); err != nil {
		return fmt.Errorf("server run err: %v", err)
	}

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	if err := s.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("server stop err: %v", err)
	}

	return nil
}

func (s *Server) registerUser() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		//vadation json...

		//run app
		_, err := s.usecase.RegisterUser(ctx, &application.InputRegisterUser{Email: "hello", Password: "world"})
		if err != nil {
			switch err.(type) {
			case *application.UserAlreadyExistError:
				w.WriteHeader(http.StatusConflict)
				return
			default:
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}

		//encode json...
		w.WriteHeader(200)
	}
}

func (s *Server) listItems() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		//vadation json...

		//run app
		_, err := s.usecase.ListItems(ctx)
		if err != nil {
			switch err.(type) {
			default:
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}

		//encode json...
		//write body...
		w.WriteHeader(200)
	}
}

func (s *Server) setupPayment() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		//vadation json...

		//run app
		_, err := s.usecase.SetupPayment(ctx, &application.InputSetupPayment{UserID: "hello", CreditCardNumber: "world"})
		if err != nil {
			switch err.(type) {
			case *application.PaymentAlreadySetupError:
				w.WriteHeader(http.StatusConflict)
				return
			default:
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}

		//encode json...
		//write body...
		w.WriteHeader(200)
	}
}

func (s *Server) buyItem() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		//vadation json...

		//run app
		err := s.usecase.BuyItem(ctx, &application.InputBuyItem{UserID: "hello", ItemID: "world"})
		if err != nil {
			switch err.(type) {
			case *application.NotPaymentAccountRegisteredError:
				w.WriteHeader(http.StatusBadRequest)
				return
			case *application.ItemIsAlreadySoldError:
				w.WriteHeader(http.StatusConflict)
				return
			default:
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}

		//encode json...
		//write body...
		w.WriteHeader(200)
	}
}
