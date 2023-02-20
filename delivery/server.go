package delivery

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"simple-payment/delivery/controller"
	_ "simple-payment/docs"
)

type Server struct {
	engine *gin.Engine
	host string
}

func (s *Server) initController() {
	publicRoute := s.engine.Group("/api")
	controller.NewCustomerController(publicRoute)
	controller.NewMerchantController(publicRoute)
	controller.NewBankController(publicRoute)
}

func (s *Server) Run() {
	s.initController()

	s.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
