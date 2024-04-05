package service_test

import (
	"course-sign-up/internal/service"
	"course-sign-up/pkg/config"
	"course-sign-up/pkg/log"
	mock_repository "course-sign-up/test/mocks/repository"

	"context"
	"flag"
	"fmt"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var (
	logger *log.Logger
)

func TestMain(m *testing.M) {
	fmt.Println("begin")

	err := os.Setenv("APP_CONF", "../../../config/local.yml")
	if err != nil {
		panic(err)
	}

	var envConf = flag.String("conf", "config/local.yml", "config path, eg: -conf ./config/local.yml")
	flag.Parse()
	conf := config.NewConfig(*envConf)

	logger = log.NewLog(conf)

	code := m.Run()
	fmt.Println("test end")

	os.Exit(code)
}

func TestUserService_GetSignedUpCourses(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCourseRepo := mock_repository.NewMockCourseRepository(ctrl)
	mockEnrollmentRepo := mock_repository.NewMockEnrollmentRepository(ctrl)
	srv := service.NewService(logger)

	courseService := service.NewCourseService(srv, mockCourseRepo, mockEnrollmentRepo)

	ctx := context.Background()

	mockCourseRepo.EXPECT().ListSignedUpCourses(ctx, "test@mail.com").Return(nil, nil)

	_, err := courseService.GetSignedUpCourses(ctx, "test@mail.com")

	assert.NoError(t, err)
}
