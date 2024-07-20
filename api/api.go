package api

import (
	"fmt"

	"github.com/bhyago/api-crud-go/infra/config"
	"github.com/gin-gonic/gin"
)

type Service struct {
	*gin.Engine
}

func NewService() *Service {
	r := gin.Default()
	return &Service{r}
}

func (s *Service) Start() error {
	s.GetRoutes()
	return s.Engine.Run(fmt.Sprintf(":%d", config.Env.ApiPort))
}
