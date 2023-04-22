package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"project-structure/pkg/common"
	"project-structure/pkg/logger"
)

const RESTATUSABORT = 204

// API is the interface for the API server.
type API interface {
	// Start starts the API server
	Start() error
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

// apiServer is the implementation of the API interface for gin.
type apiServer struct {
	port     int
	gin      *gin.Engine
	lg       *logger.Entry
	services *common.Services
}

// Start starts the API server.
func (s *apiServer) Start() error {
	return s.gin.Run(fmt.Sprintf(":%d", s.port))
}

type ServerConfig struct {
	Port     int
	GinMode  string
	Services *common.Services
}

// NewServer creates a new API server.
func NewServer(c *ServerConfig) API {
	gin.SetMode(c.GinMode)
	ginHandler := gin.Default()

	apiLogger := logger.General.Component("api")

	server := &apiServer{
		port:     c.Port,
		gin:      ginHandler,
		lg:       apiLogger,
		services: c.Services,
	}

	server.AddCors()
	server.RegisterHealthCheck()
	server.RegisterGraphql()
	server.RegisterWebhooks()

	apiLogger.Info("API server started")

	return server
}

func (s *apiServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.gin.ServeHTTP(w, r)
}

// RegisterHealthCheck registers a health check endpoint.
func (s *apiServer) RegisterHealthCheck() {
	s.gin.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
}

func (s *apiServer) AddCors() {
	s.gin.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT, DELETE")
		c.Next()
	})
}
