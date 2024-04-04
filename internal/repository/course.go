package repository

import (
	"context"
	"course-sign-up/internal/model"
)

type CourseRepository interface {
	FirstById(ctx context.Context, id int64) (*model.Course, error)
	List(ctx context.Context) ([]*model.Course, error)
}

func NewCourseRepository(repository *Repository) CourseRepository {
	return &courseRepository{
		Repository: repository,
	}
}

type courseRepository struct {
	*Repository
}

func (r *courseRepository) FirstById(ctx context.Context, id int64) (*model.Course, error) {
	var course model.Course
	// TODO: query db
	return &course, nil
}

func (r *courseRepository) List(ctx context.Context) ([]*model.Course, error) {
	var courses []*model.Course
	// Query the database for all courses. The context is passed to allow query cancelation.
	result := r.db.WithContext(ctx).Find(&courses)

	if result.Error != nil { // Handling the possible errors that might occur during the query.
		return nil, result.Error
	}
	return courses, nil
}
