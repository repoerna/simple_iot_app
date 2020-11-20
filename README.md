# The Idea
is to implement something like this:
[![N|Solid](https://dzone.com/storage/temp/1788474-1444132453algorithm-basic.png)](https://dzone.com/articles/enabling-caching-in-mongodb-database-with-redis-us)

# Technology
- Go
- Gorilla/Mux for router
- Postgres
- Redis
- SQLC for generate queries

# Project Structure
- **/db**: as data layer, contain queries and migration files
- **/api/config**: project configuration (db, redis, server)
- **/api/handler**: as business layer, parsing and validate request
- **/api/**: cotain server file to init db, redis, and router
- **/Makefile**: for preparation

# Redis Implementation
redis used as middleware, so we can easily integrate it using ***gorilla/mux***
- if the request GET it will check redis, if data not available, it will call db and then set cache in redis
- if the request PUT or to update, data will be process to db and update data in redis
- 
### Redis Middleware
```
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

```

### Implement in router
```
// Initialize ...
func (s *Server) Initialize(username, password, host, port, dbName, cacheAddr, cachePass string) {
	....

	s.Router.Use(
		// s.Middleware.authMiddleware,
		s.Middleware.CacheMiddleware,
	)

	s.InitializeRoutes()

}
```
# Run the App
### 1. start postgre server
```sh
$ make postgre
```
### 2. start redis server
```sh
$ make redis
```
### 3. create database
```sh
$ make createdb
```
### 4. migrate table
```sh
$ make migrateup
```
### 5. start application
```sh
$ make server
```

### 6. db test
```sh
$ make test
```

## Test the Endpoint
- GET /api/devices/{id}
- GET /api/devices
- POST /api/devices
- PUT /api/devices/{id}
- DELETE /api/devices/{id}