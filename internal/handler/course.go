package handler

import (
	v1 "course-sign-up/api/v1"
	"course-sign-up/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CourseHandler interface {
	ListCourses(ctx *gin.Context)
	SignUpCourse(ctx *gin.Context)
	GetSignedUpCourse(ctx *gin.Context)
	DeleteSignedUpCourse(ctx *gin.Context)
	GetCourseClassmates(ctx *gin.Context)
}

type courseHandler struct {
	*Handler
	courseService service.CourseService
}

func NewCourseHandler(handler *Handler, courseService service.CourseService) *courseHandler {
	return &courseHandler{
		Handler:       handler,
		courseService: courseService,
	}
}

func (h *courseHandler) ListCourses(ctx *gin.Context) {
	courses, err := h.courseService.ListCourses(ctx)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, v1.ErrInternalServerError, nil)
		return
	}

	v1.HandleSuccess(ctx, courses)
}

func (h *courseHandler) SignUpCourse(ctx *gin.Context) {

}

func (h *courseHandler) GetSignedUpCourse(ctx *gin.Context) {

}
func (h *courseHandler) DeleteSignedUpCourse(ctx *gin.Context) {

}

func (h *courseHandler) GetCourseClassmates(ctx *gin.Context) {

}
