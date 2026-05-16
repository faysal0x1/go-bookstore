package middleware

import (
	"net/http"

	"github.com/didip/tollbooth/v7"
	"github.com/didip/tollbooth/v7/limiter"
)

func RateLimitMiddleware() func(http.Handler) http.Handler {
	// Create a limiter that allows 5 requests per second
	lmt := tollbooth.NewLimiter(5, &limiter.ExpirableOptions{DefaultExpirationTTL: 3600})
	lmt.SetMessage("You have reached the maximum number of requests. Please try again later.")
	
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			httpError := tollbooth.LimitByRequest(lmt, w, r)
			if httpError != nil {
				// tollbooth already writes the error to the response
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
