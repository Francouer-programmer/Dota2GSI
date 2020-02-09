package app

import (

	"github.com/go-programmer/Dota2GSI/app/handler"
)

//Run application
func Run() error {
	if err := handler.Run(); err != nil {
		return err
	}
	return nil 
}


