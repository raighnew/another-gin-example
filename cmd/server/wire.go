//go:build wireinject
// +build wireinject

package main

import (
	"course-sign-up/internal/handler"
	"course-sign-up/internal/repository"
	"course-sign-up/internal/server"
	"course-sign-up/internal/service"
	"course-sign-up/pkg/log"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

var ServerSet = wire.NewSet(server.NewServerHTTP)

var RepositorySet = wire.NewSet(
	repository.NewDb,
	repository.NewRepository,
	repository.NewCourseRepository,
)

var ServiceSet = wire.NewSet(
	service.NewService,
	service.NewCourseService,
)

var HandlerSet = wire.NewSet(
	handler.NewHandler,
	handler.NewCourseHandler,
)

func newApp(*viper.Viper, *log.Logger) (*gin.Engine, func(), error) {
	panic(wire.Build(
		ServerSet,
		RepositorySet,
		ServiceSet,
		HandlerSet,
	))
}
