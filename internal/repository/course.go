package repository

import (
	"context"
	"course-sign-up/internal/model"
)

type CourseRepository interface {
	List(ctx context.Context) ([]*model.Course, error)
	Exists(ctx context.Context, courseID string) (bool, error)
	ListSignedUpCourses(ctx context.Context, studentEmail string) ([]*model.Course, error)
}

func NewCourseRepository(repository *Repository) CourseRepository {
	return &courseRepository{
		Repository: repository,
	}
}

type courseRepository struct {
	*Repository
}

func (r *courseRepository) List(ctx context.Context) ([]*model.Course, error) {
	var courses []*model.Course

	result := r.db.WithContext(ctx).Find(&courses)

	if result.Error != nil {
		return nil, result.Error
	}
	return courses, nil
}

func (r *courseRepository) ListSignedUpCourses(ctx context.Context, studentEmail string) ([]*model.Course, error) {
	var courses []*model.Course

	result := r.db.WithContext(ctx).Table("courses").
		Joins("RIGHT JOIN enrollments ON courses.course_id = enrollments.course_id").
		Where("enrollments.student_id = ?", studentEmail).
		Find(&courses)

	if result.Error != nil {
		return nil, result.Error
	}

	return courses, nil
}

func (r *courseRepository) Exists(ctx context.Context, courseID string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&model.Course{}).Where("course_id = ?", courseID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
