package delivery

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"simple-payment/config"
	"simple-payment/delivery/controller"
	"simple-payment/delivery/middleware"
	_ "simple-payment/docs"
	"simple-payment/manager"
	"simple-payment/util/authenthicator"
)

type Server struct {
	engine         *gin.Engine
	host           string
	useCaseManager manager.UseCaseManager
	tokenService   authenthicator.AccessToken
}

func (s *Server) initController() {
	publicRoute := s.engine.Group("/api")
	tokenMdw := middleware.NewTokenValidator(s.tokenService)
	controller.NewCustomerController(publicRoute, s.useCaseManager.CustomerUseCase(), tokenMdw)
	controller.NewMerchantController(publicRoute, s.useCaseManager.MerchantUseCase())
	controller.NewBankController(publicRoute, s.useCaseManager.BankUseCase())
	controller.NewPaymentController(publicRoute, s.useCaseManager.PaymentUseCase())
	controller.NewUserController(publicRoute, s.useCaseManager.UserUseCase())
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

	conf := config.NewConfig()
	infra := manager.NewInfraManager(conf)
	tokenService := authenthicator.NewAccessToken(conf)

	rm := manager.NewRepositoryManager(infra)
	ucm := manager.NewUseCaseManager(rm, tokenService)

	return &Server{
		engine:         router,
		host:           host,
		useCaseManager: ucm,
		tokenService:   tokenService,
	}
}
