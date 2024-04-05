package repository

import (
	"context"
	v1 "course-sign-up/api/v1"
	"course-sign-up/internal/model"
)

type EnrollmentRepository interface {
	CreateEnrollment(ctx context.Context, studentId string, courseId string) (*model.Enrollment, error)
	DeleteEnrollment(ctx context.Context, studentId string, courseId string) (bool, error)
	GetEnrollment(ctx context.Context, studentId string, courseId string) (*model.Enrollment, error)
	GetCourseClassmates(ctx context.Context, studentId string, courseId string) ([]*model.Enrollment, error)
}

func NewEnrollmentRepository(repository *Repository) EnrollmentRepository {
	return &enrollmentRepository{
		Repository: repository,
	}
}

type enrollmentRepository struct {
	*Repository
}

func (r *enrollmentRepository) CreateEnrollment(ctx context.Context, studentId string, courseId string) (*model.Enrollment, error) {
	enrollment := &model.Enrollment{
		StudentID: studentId,
		CourseID:  courseId,
	}

	err := r.db.WithContext(ctx).Create(enrollment).Error

	if err != nil {
		return nil, err
	}

	return enrollment, nil
}

func (r *enrollmentRepository) DeleteEnrollment(ctx context.Context, studentId string, courseId string) (bool, error) {
	var enrollment *model.Enrollment

	result := r.db.WithContext(ctx).Where("student_id = ? AND course_id = ?", studentId, courseId).Delete(enrollment)

	if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}

func (r *enrollmentRepository) GetEnrollment(ctx context.Context, studentId string, courseId string) (*model.Enrollment, error) {
	var enrollment *model.Enrollment

	result := r.db.WithContext(ctx).Where("student_id = ? AND course_id = ?", studentId, courseId).Find(&enrollment)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, v1.ErrNotFound
	}

	return enrollment, nil
}

func (r *enrollmentRepository) GetCourseClassmates(ctx context.Context, studentId string, courseId string) ([]*model.Enrollment, error) {
	var enrollments []*model.Enrollment

	result := r.db.WithContext(ctx).Where("student_id != ? AND course_id = ?", studentId, courseId).Find(&enrollments)

	if result.Error != nil {
		return nil, result.Error
	}

	return enrollments, nil
}
