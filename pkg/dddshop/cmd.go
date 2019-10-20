package dddshop

import (
	"fmt"

	"github.com/sueken5/golang-ddd/pkg/dddshop/interfaces/http"
)

func Execute() error {
	//di...

	srv := http.NewServer()
	if err := srv.Run(); err != nil {
		fmt.Errorf("dddshop exec err: %v", err)
	}

	return nil
}
