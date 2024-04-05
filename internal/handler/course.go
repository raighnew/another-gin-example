package handler

import (
	v1 "course-sign-up/api/v1"
	"course-sign-up/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CourseHandler struct {
	*Handler
	courseService service.CourseService
}

func NewCourseHandler(handler *Handler, courseService service.CourseService) *CourseHandler {
	return &CourseHandler{
		Handler:       handler,
		courseService: courseService,
	}
}

func (h *CourseHandler) ListCourses(ctx *gin.Context) {
	courses, err := h.courseService.ListCourses(ctx)

	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, v1.ErrInternalServerError, nil)
		return
	}

	v1.HandleSuccess(ctx, courses)
}

func (h *CourseHandler) SignUpCourse(ctx *gin.Context) {
	studentEmail := ctx.Param("studentEmail")
	req := new(v1.SignUpRequest)

	if err := ctx.ShouldBindJSON(req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}

	// check course exists
	exists, err := h.courseService.IfCourseExists(ctx, req.CourseID)
	if err != nil {
		// Handle the error, could be ErrInternal or something similar
		v1.HandleError(ctx, http.StatusInternalServerError, v1.ErrInternalServerError, err.Error())
		return
	}
	if !exists {
		v1.HandleError(ctx, http.StatusNotFound, v1.ErrNotFound, "Course not found.")
		return
	}
	// check course already signed up
	enrollment, err := h.courseService.SignUpCourse(ctx, studentEmail, req.CourseID)

	if err != nil {
		v1.HandleError(ctx, http.StatusConflict, v1.ErrConflict, "Student has already signed up for this course.")
		return
	}

	v1.HandleSuccess(ctx, enrollment)
}

func (h *CourseHandler) GetSignedUpCourses(ctx *gin.Context) {
	studentEmail := ctx.Param("studentEmail")

	enrollments, err := h.courseService.GetSignedUpCourses(ctx, studentEmail)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, v1.ErrInternalServerError, "Error retrieving signed up courses")
		return
	}
	v1.HandleSuccess(ctx, enrollments)
}

func (h *CourseHandler) DeleteSignedUpCourse(ctx *gin.Context) {
	studentEmail := ctx.Param("studentEmail")
	courseId := ctx.Param("courseId")

	if courseId == "" || studentEmail == "" {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, "Error parameters")
		return
	}

	enrollments, err := h.courseService.DeleteSignedUpCourse(ctx, studentEmail, courseId)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, v1.ErrInternalServerError, "Error deleting signed up courses")
		return
	}
	v1.HandleSuccess(ctx, enrollments)

}

func (h *CourseHandler) GetCourseClassmates(ctx *gin.Context) {
	studentEmail := ctx.Param("studentEmail")
	courseId := ctx.Param("courseId")

	if courseId == "" || studentEmail == "" {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, "Error parameters")
		return
	}
	//  Check student already signed for this course
	_, err := h.courseService.GetCourseEnrollment(ctx, studentEmail, courseId)

	if err != nil {
		v1.HandleError(ctx, http.StatusNotFound, v1.ErrNotFound, "Student is not signed up for this course.")
		return
	}

	enrollments, err := h.courseService.GetCourseClassmates(ctx, studentEmail, courseId)

	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, v1.ErrInternalServerError, "Error retrieving signed up courses")
		return
	}

	v1.HandleSuccess(ctx, enrollments)
}
