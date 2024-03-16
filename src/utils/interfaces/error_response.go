package utils_interfaces

type ErrorResponse struct {
	Error      string `json:"error"`
	StatusCode int    `json:"statusCode"`
}
