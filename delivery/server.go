package delivery

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"simple-payment/delivery/controller"
)

type Server struct {
	engine *gin.Engine
	host string
}

func (s *Server) initController() {
	publicRoute := s.engine.Group("api")
	controller.NewCustomerController(publicRoute)
}

func (s *Server) Run() {
	s.initController()

	if err := s.engine.Run(s.host); err != nil {
		panic(err)
	} else {
		fmt.Println("Server is running on port", s.host)
	}
}

func NewServer(host string) *Server {
	router := gin.Default()

	return &Server{
		engine: router,
		host: host,
	}
}
