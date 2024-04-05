package v1

// Define request body here
type SignUpRequest struct {
	CourseID string `json:"courseId" binding:"required" example:"CS101"`
}
