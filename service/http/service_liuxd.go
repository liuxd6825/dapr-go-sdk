package http

import (
	"github.com/dapr/go-sdk/service/common"
	"github.com/dapr/go-sdk/service/internal"
	"github.com/kataras/iris/v12"
	"net/http"
	"os"
)

// NewService creates new Service.
func NewIrisService(app *iris.Application, httpServer *http.Server) common.Service {
	return newIrisServer(app, httpServer)
}

func newIrisServer(app *iris.Application, httpServer *http.Server) *Server {
	return &Server{
		address:        httpServer.Addr,
		httpServer:     httpServer,
		mux:            NewIrisMux(app),
		topicRegistrar: make(internal.TopicRegistrar),
		authToken:      os.Getenv(common.AppAPITokenEnvVar),
	}
}

func (s *Server) RegisterBaseHandler() {
	s.registerBaseHandler()
}

// Start starts the HTTP handler. Blocks while serving.
func (s *Server) Start() error {
	s.registerBaseHandler()
	return nil
	//return s.httpServer.ListenAndServe()
}
