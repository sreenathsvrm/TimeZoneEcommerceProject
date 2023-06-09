//go:build wireinject
// +build wireinject

package di

import (
	http "ecommerce/pkg/api"
	handler "ecommerce/pkg/api/handler"
	config "ecommerce/pkg/config"
	db "ecommerce/pkg/db"
	repository "ecommerce/pkg/repository"
	usecase "ecommerce/pkg/usecase"

	"github.com/google/wire"
)

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(db.ConnectDatabase,
		repository.NewUserRepository,
		repository.NewAdminRepository,
		repository.NewproductRepository,
		repository.NewecartRepository,
		repository.NewOrderRepository,
		repository.NewCouponrepo,
		usecase.NewUserUseCase,
		usecase.NewOtpUseCase,
		usecase.NewAdminUseCase,
		usecase.NewProductUsecase,
		usecase.NewCartUsecase,
		usecase.NewOrderUseCase,
        usecase.NewCouponUseCase,
		handler.NewUserHandler,
		handler.NewOtpHandler,
		handler.NewAdminHandler,
		handler.NewproductHandler,
        handler.NewCartHandler,
		handler.NewOrderHandler,
		handler.NewCoupenHandler,
		http.NewServerHTTP,
		
	)

	return &http.ServerHTTP{}, nil
}
