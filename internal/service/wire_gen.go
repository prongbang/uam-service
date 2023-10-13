// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package service

import (
	"github.com/casbin/casbin/v2"
	"github.com/prongbang/uam-service/internal/service/database"
	"github.com/prongbang/uam-service/internal/service/uam"
	"github.com/prongbang/uam-service/internal/shared/auth"
	"github.com/prongbang/uam-service/internal/shared/role"
	"github.com/prongbang/uam-service/internal/shared/user"
)

// Injectors from wire.go:

func New(dbDriver database.Drivers, enforce *casbin.Enforcer) Service {
	dataSource := user.NewDataSource(dbDriver)
	repository := user.NewRepository(dataSource)
	useCase := user.NewUseCase(repository, enforce)
	roleDataSource := role.NewDataSource(dbDriver)
	roleRepository := role.NewRepository(roleDataSource)
	roleUseCase := role.NewUseCase(roleRepository)
	handler := role.NewHandler(useCase, roleUseCase)
	authDataSource := auth.NewDataSource(dbDriver)
	authRepository := auth.NewRepository(authDataSource)
	authUseCase := auth.NewUseCase(authRepository, roleUseCase)
	authHandler := auth.NewHandler(authUseCase)
	userHandler := user.NewHandler(useCase)
	validate := role.NewValidate()
	authValidate := auth.NewValidate()
	userValidate := user.NewValidate()
	apiRouter := uam.NewRouter(handler, authHandler, userHandler, validate, authValidate, userValidate)
	serviceRouters := NewRouters(apiRouter)
	serviceAPI := NewAPI(serviceRouters)
	uamServer := uam.NewServer(useCase, authUseCase)
	grpcListener := uam.NewListener(uamServer)
	grpc := NewGRPC(grpcListener)
	serviceService := NewService(serviceAPI, grpc)
	return serviceService
}
