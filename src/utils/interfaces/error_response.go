package utils_interfaces

type ErrorResponse struct {
	Error      string `json:"error" example:"variable error message"`
	StatusCode int    `json:"statusCode" example:"-1"`
}
