package delivery

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
	host string
}

func (s *Server) Run() {
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
