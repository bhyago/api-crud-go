package main

import (
	"github.com/bhyago/api-crud-go/api"
)

func main() {
	service := api.NewService()
	service.Start()
}
