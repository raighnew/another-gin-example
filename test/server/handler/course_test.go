package handler

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"

	v1 "course-sign-up/api/v1"
	"course-sign-up/internal/handler"
	"course-sign-up/internal/middleware"
	"course-sign-up/pkg/config"
	"course-sign-up/pkg/log"

	"net/http"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var logger *log.Logger
var router *gin.Engine

func TestMain(m *testing.M) {
	fmt.Println("begin")
	err := os.Setenv("APP_CONF", "../../../config/local.yml")
	if err != nil {
		fmt.Println("Setenv error", err)
	}
	var envConf = flag.String("conf", "config/local.yml", "config path, eg: -conf ./config/local.yml")
	flag.Parse()
	conf := config.NewConfig(*envConf)

	logger = log.NewLog(conf)
	hdl = handler.NewHandler(logger)

	gin.SetMode(gin.TestMode)
	router = gin.Default()
	router.Use(
		middleware.ResponseLogMiddleware(logger),
		middleware.RequestLogMiddleware(logger),
		//middleware.SignMiddleware(log),
	)

	code := m.Run()
	fmt.Println("test end")

	os.Exit(code)
}

func TestUserHandler_Register(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	params := v1.RegisterRequest{
		Password: "123456",
		Email:    "xxx@gmail.com",
	}

	mockUserService := mock_service.NewMockUserService(ctrl)
	mockUserService.EXPECT().Register(gomock.Any(), &params).Return(nil)

	userHandler := handler.NewUserHandler(hdl, mockUserService)
	router.POST("/register", userHandler.Register)

	paramsJson, _ := json.Marshal(params)

	resp := performRequest(router, "POST", "/register", bytes.NewBuffer(paramsJson))

	assert.Equal(t, resp.Code, http.StatusOK)
	// Add assertions for the response body if needed
}
