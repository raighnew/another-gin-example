package service

import (
	"context"
	"course-sign-up/internal/model"
	"course-sign-up/internal/repository"
)

type CourseService interface {
	ListCourses(ctx context.Context) ([]*model.Course, error)
	IfCourseExists(ctx context.Context, courseId string) (bool, error)
	SignUpCourse(ctx context.Context, studentEmail string, courseId string) (*model.Enrollment, error)
	GetSignedUpCourses(ctx context.Context, studentEmail string) ([]*model.Course, error)
	DeleteSignedUpCourse(ctx context.Context, studentEmail string, courseId string) (bool, error)
	GetCourseClassmates(ctx context.Context, studentEmail string, courseId string) ([]*model.Enrollment, error)
	GetCourseEnrollment(ctx context.Context, studentEmail string, courseId string) (*model.Enrollment, error)
}

func NewCourseService(service *Service, courseRepository repository.CourseRepository, enrollmentRepository repository.EnrollmentRepository) CourseService {
	return &courseService{
		Service:              service,
		courseRepository:     courseRepository,
		enrollmentRepository: enrollmentRepository,
	}
}

type courseService struct {
	courseRepository     repository.CourseRepository
	enrollmentRepository repository.EnrollmentRepository
	*Service
}

func (s *courseService) ListCourses(ctx context.Context) ([]*model.Course, error) {
	return s.courseRepository.List(ctx)
}

func (s *courseService) IfCourseExists(ctx context.Context, courseId string) (bool, error) {
	return s.courseRepository.Exists(ctx, courseId)
}

func (s *courseService) SignUpCourse(ctx context.Context, studentEmail string, courseId string) (*model.Enrollment, error) {
	return s.enrollmentRepository.CreateEnrollment(ctx, studentEmail, courseId)
}

func (s *courseService) GetSignedUpCourses(ctx context.Context, studentEmail string) ([]*model.Course, error) {
	return s.courseRepository.ListSignedUpCourses(ctx, studentEmail)
}

func (s *courseService) DeleteSignedUpCourse(ctx context.Context, studentEmail string, courseId string) (bool, error) {
	return s.enrollmentRepository.DeleteEnrollment(ctx, studentEmail, courseId)
}

func (s *courseService) GetCourseEnrollment(ctx context.Context, studentEmail string, courseId string) (*model.Enrollment, error) {
	return s.enrollmentRepository.GetEnrollment(ctx, studentEmail, courseId)
}

func (s *courseService) GetCourseClassmates(ctx context.Context, studentEmail string, courseId string) ([]*model.Enrollment, error) {
	return s.enrollmentRepository.GetCourseClassmates(ctx, studentEmail, courseId)
}
