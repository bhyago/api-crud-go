package main

import (
	"log"

	"github.com/bhyago/api-crud-go/api"
	"github.com/bhyago/api-crud-go/infra/config"
)

func main() {
	var err error

	err = config.StartConfig()
	FatalError(err)

	err = api.NewService().Start()
	FatalError(err)
}

func FatalError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
