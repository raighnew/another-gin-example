package service

import (
	"context"
	"course-sign-up/internal/model"
	"course-sign-up/internal/repository"
)

type CourseService interface {
	ListCourses(ctx context.Context) ([]*model.Course, error)
	SignUpCourse(ctx context.Context, id int64) (*model.Course, error)
	GetSignedUpCourse(ctx context.Context, id int64) (*model.Course, error)
	DeleteSignedUpCourse(ctx context.Context, id int64) (*model.Course, error)
	GetCourseClassmates(ctx context.Context, id int64) (*model.Course, error)
}

func NewCourseService(service *Service, courseRepository repository.CourseRepository) CourseService {
	return &courseService{
		Service:          service,
		courseRepository: courseRepository,
	}
}

type courseService struct {
	*Service
	courseRepository repository.CourseRepository
}

func (s *courseService) ListCourses(ctx context.Context) ([]*model.Course, error) {
	return s.courseRepository.List(ctx)
}

func (s *courseService) SignUpCourse(ctx context.Context, id int64) (*model.Course, error) {
	return s.courseRepository.FirstById(ctx, id)
}

func (s *courseService) GetSignedUpCourse(ctx context.Context, id int64) (*model.Course, error) {
	return s.courseRepository.FirstById(ctx, id)
}

func (s *courseService) DeleteSignedUpCourse(ctx context.Context, id int64) (*model.Course, error) {
	return s.courseRepository.FirstById(ctx, id)
}

func (s *courseService) GetCourseClassmates(ctx context.Context, id int64) (*model.Course, error) {
	return s.courseRepository.FirstById(ctx, id)
}
