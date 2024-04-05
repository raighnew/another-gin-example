package server

import (
	"course-sign-up/internal/handler"
	"course-sign-up/pkg/helper/resp"
	"course-sign-up/pkg/log"

	"github.com/gin-gonic/gin"
)

func NewServerHTTP(
	logger *log.Logger,
	courseHandler *handler.CourseHandler,
) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/health", func(ctx *gin.Context) {
		resp.HandleSuccess(ctx, map[string]interface{}{
			"health": "Okay",
		})
	})

	r.GET("/courses", courseHandler.ListCourses)
	r.POST("/students/:studentEmail/courses", courseHandler.SignUpCourse)
	r.GET("/students/:studentEmail/courses", courseHandler.GetSignedUpCourses)
	r.DELETE("/students/:studentEmail/courses/:courseId", courseHandler.DeleteSignedUpCourse)
	r.GET("/students/:studentEmail/courses/:courseId/classmates", courseHandler.GetCourseClassmates)

	return r
}
