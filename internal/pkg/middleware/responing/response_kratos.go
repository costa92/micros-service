package responding

import (
	"context"
	"net/http"
	"time"

	"github.com/go-kratos/kratos/v2/middleware"
)

type response struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Ts      string      `json:"ts"`
}

// StandardizedResponseMiddleware wraps responses in a standardized JSON format.
func StandardizedResponseMiddleware() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			// Call the next handler in the chain
			resp, err := handler(ctx, req)
			if err != nil {
				// Use helper function to build error response
				return buildResponse(http.StatusInternalServerError, nil, err.Error()), nil
			}

			// Check if the response includes a message
			var message string
			if msg, ok := resp.(map[string]interface{})["message"]; ok {
				message = msg.(string)
			} else {
				message = "success"
			}

			// Use helper function to build successful response
			return buildResponse(http.StatusOK, resp, message), nil
		}
	}
}

// buildResponse constructs a standardized response.
func buildResponse(code int, data interface{}, message string) *response {
	return &response{
		Code:    code,
		Data:    data,
		Message: message,
		Ts:      time.Now().UTC().Format(time.RFC3339), // Use UTC for consistency
	}
}
