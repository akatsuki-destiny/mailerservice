package services

import "github.com/gofiber/fiber/v2"

type ErrorResponse struct {
	Status       string `json:"status"`
	Error        error  `json:"error"`
	ErrorMessage string `json:"error_message"`
}

type Response struct {
	Status string                 `json:"status"`
	Data   map[string]interface{} `json:"data"`
}

// HandleErrorResponse handles error response
func HandleErrorResponse(c *fiber.Ctx, statusCode int, err error) error {

	if err != nil {

		resp := ErrorResponse{
			Status:       "failed",
			Error:        err,
			ErrorMessage: err.Error(),
		}

		return c.Status(statusCode).JSON(resp)
	}
	return nil
}
