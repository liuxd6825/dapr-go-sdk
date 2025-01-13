package http

import (
	"github.com/dapr/go-sdk/service/common"
	"github.com/dapr/go-sdk/service/internal"
	"github.com/kataras/iris/v12"
	"net/http"
	"os"
)

// NewMyService creates new Service.
func NewMyService(app *iris.Application, httpServer *http.Server) common.Service {
	return newMyServer(app, httpServer)
}

func newMyServer(app *iris.Application, httpServer *http.Server) *Server {
	return &Server{
		httpServer:     httpServer,
		mux:            NewIrisMux(app),
		topicRegistrar: make(internal.TopicRegistrar),
		authToken:      os.Getenv(common.AppAPITokenEnvVar),
	}
}

func (s *Server) RegisterBaseHandler() {
	s.registerBaseHandler()
}
