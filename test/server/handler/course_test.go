package handler_test

import (
	"bytes"
	v1 "course-sign-up/api/v1"
	"course-sign-up/internal/handler"
	"course-sign-up/internal/model"
	"course-sign-up/pkg/config"
	"course-sign-up/pkg/log"
	mock_service "course-sign-up/test/mocks/service"
	"encoding/json"

	"flag"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var logger *log.Logger
var router *gin.Engine
var hdl *handler.Handler

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

	code := m.Run()
	fmt.Println("test end")

	os.Exit(code)
}

func TestCourseHandler_List(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	courseService := mock_service.NewMockCourseService(ctrl)
	courseService.EXPECT().ListCourses(gomock.Any()).Return(nil, nil)

	coursesHandler := handler.NewCourseHandler(hdl, courseService)

	router.GET("/courses", coursesHandler.ListCourses)

	req, _ := http.NewRequest("GET", "/courses", nil)

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)
	assert.Equal(t, resp.Code, http.StatusOK)
}

func TestCourseHandler_SignUp(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	params := v1.SignUpRequest{
		CourseID: "CS101",
	}

	courseService := mock_service.NewMockCourseService(ctrl)
	courseService.EXPECT().IfCourseExists(gomock.Any(), "CS101").Return(true, nil)
	courseService.EXPECT().SignUpCourse(gomock.Any(), gomock.Any(), "CS101").Return(&model.Enrollment{
		StudentID: "test@mail.com",
		CourseID:  "CS101",
	}, nil)

	coursesHandler := handler.NewCourseHandler(hdl, courseService)

	router.POST("/students/:studentEmail/courses", coursesHandler.SignUpCourse)

	paramsJson, _ := json.Marshal(params)

	req, _ := http.NewRequest("POST", "/students/test@mail.com/courses", bytes.NewBuffer(paramsJson))

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)
	assert.Equal(t, resp.Code, http.StatusOK)
}
