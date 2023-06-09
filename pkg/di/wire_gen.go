// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"ecommerce/pkg/api"
	"ecommerce/pkg/api/handler"
	"ecommerce/pkg/config"
	"ecommerce/pkg/db"
	"ecommerce/pkg/repository"
	"ecommerce/pkg/usecase"
)

// Injectors from wire.go:

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}
	userRepository := repository.NewUserRepository(gormDB)
	userUseCase := usecase.NewUserUseCase(userRepository)
	userHandler := handler.NewUserHandler(userUseCase)
	otpUseCase := usecase.NewOtpUseCase(cfg)
	otpHandler := handler.NewOtpHandler(cfg, otpUseCase, userUseCase)
	adminRepository := repository.NewAdminRepository(gormDB)
	adminUsecase := usecase.NewAdminUseCase(adminRepository)
	adminHandler := handler.NewAdminHandler(adminUsecase)
	productRepo := repository.NewproductRepository(gormDB)
	productUsecase := usecase.NewProductUsecase(productRepo)
	productHandler := handler.NewproductHandler(productUsecase)
	cartRepo := repository.NewecartRepository(gormDB)
	cartUsecase := usecase.NewCartUsecase(cartRepo)
	cartHandler := handler.NewCartHandler(cartUsecase)
	orderRepo := repository.NewOrderRepository(gormDB)
	orderusecase := usecase.NewOrderUseCase(orderRepo, cartRepo)
	orderHandler := handler.NewOrderHandler(orderusecase)
	couponRepo := repository.NewCouponrepo(gormDB)
	couponUseCase := usecase.NewCouponUseCase(couponRepo)
	couponHandler := handler.NewCoupenHandler(couponUseCase)
	serverHTTP := http.NewServerHTTP(userHandler, otpHandler, adminHandler, productHandler, cartHandler, orderHandler, couponHandler)
	return serverHTTP, nil
}
