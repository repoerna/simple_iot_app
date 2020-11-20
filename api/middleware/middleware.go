package middleware

import (
	"net/http"
	"net/http/httptest"
	"simpleiotapp/util"
	"time"

	"github.com/go-redis/redis"
)

// Middleware ...
type Middleware struct {
	Cache *redis.Client
}

// CacheMiddleware ...
func (m *Middleware) CacheMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// cache will be use only when get data or update data
		if r.Method == "GET" {
			content, err := m.Cache.Get(r.RequestURI).Result()
			if err != nil {
				rr := httptest.NewRecorder()
				next.ServeHTTP(rr, r)
				content = rr.Body.String()
				err = m.Cache.Set(r.RequestURI, content, 10*time.Minute).Err()
				if err != nil {
					util.RespondWithError(w, http.StatusInternalServerError, err.Error())
				}
				util.RespondWithString(w, http.StatusOK, content)
				return
			}

		} else if r.Method == "PUT" {
			rr := httptest.NewRecorder()
			next.ServeHTTP(rr, r)
			content := rr.Body.String()
			err := m.Cache.Set(r.RequestURI, content, 10*time.Minute).Err()
			if err != nil {
				util.RespondWithError(w, http.StatusInternalServerError, err.Error())
			}
			util.RespondWithString(w, http.StatusOK, content)
			return
		}

		next.ServeHTTP(w, r)
		return
	})
}
