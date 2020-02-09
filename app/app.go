package main

import (
	"log"

	app "github.com/go-programmer/Dota2GSI/app/cmd"
)

func main() {
	ifErrorDo(app.Run())
}

func ifErrorDo(err error) {
	if err != nil {
		log.Println(err)
	}
}
