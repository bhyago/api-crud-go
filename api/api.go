package api

import "github.com/gin-gonic/gin"

type Service struct {
	*gin.Engine
}

func NewService() *Service {
	r := gin.Default()
	return &Service{r}
}

func (s *Service) Start() {
	s.GetRoutes()
	s.Engine.Run()
}
