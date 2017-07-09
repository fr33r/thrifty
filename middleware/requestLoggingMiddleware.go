package middleware

import (
	"fmt"
)

type RequestLoggingMiddleware struct{}

func (requestLoggingMiddleware *RequestLoggingMiddleware) Act() {
	fmt.Println("Log!")
}
