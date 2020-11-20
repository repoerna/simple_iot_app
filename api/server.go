package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"simpleiotapp/api/handler"
	"simpleiotapp/api/middleware"
	db "simpleiotapp/db/sqlc"

	"github.com/go-redis/redis"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

// Server ...
type Server struct {
	Router     *mux.Router
	Logger     http.Handler
	Handler    handler.Handler
	Middleware middleware.Middleware
}

// Initialize ...
func (s *Server) Initialize(username, password, host, port, dbName, cacheAddr, cachePass string) {
	dataSource := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, username, password, dbName)

	// create DB Connection
	dbConn, err := sql.Open("postgres", dataSource)
	if err != nil {
		log.Fatal(err)
	}
	// Passing connection to Queries
	s.Handler.Queries = db.New(dbConn)

	log.Printf("db connected ...")

	// create Redis connection
	cache := redis.NewClient(&redis.Options{
		Addr:     cacheAddr,
		Password: cachePass,
		DB:       0,
	})

	// Ping to redis server
	_, err = cache.Ping().Result()
	if err != nil {
		panic(err)
	}
	// passing redis connection to Cachemiddleware
	s.Middleware.Cache = cache
	log.Printf("cache db connected ...")

	// setup router using mux
	s.Router = mux.NewRouter()
	// setup logger to monitor every request
	s.Logger = handlers.CombinedLoggingHandler(os.Stdout, s.Router)

	s.Router.Use(
		// s.Middleware.authMiddleware,
		s.Middleware.CacheMiddleware,
	)

	s.InitializeRoutes()

}

// InitializeRoutes ...
func (s *Server) InitializeRoutes() {
	s.Router.HandleFunc("/api/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Test OK")
	}).Methods("GET")
	s.Router.HandleFunc("/api/devices/", s.Handler.GetDevices).Methods("GET")
	s.Router.HandleFunc("/api/devices/", s.Handler.CreateDevice).Methods("POST")
	s.Router.HandleFunc("/api/devices/{id:[0-9]+}", s.Handler.GetDevice).Methods("GET")
	s.Router.HandleFunc("/api/devices/{id:[0-9]+}", s.Handler.UpdateDevice).Methods("PUT")
	s.Router.HandleFunc("/api/devices/{id:[0-9]+}", s.Handler.DeleteDevice).Methods("DELETE")
}

// Run ...
func (s *Server) Run(addr string) {
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	log.Fatal(http.ListenAndServe(":"+viper.GetString("Server.port"),
		handlers.CORS(headersOk, originsOk, methodsOk)(s.Logger)))

	log.Printf("server is ready ...")
}
