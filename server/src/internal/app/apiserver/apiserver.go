package apiserver

import (
	// "github.com/kletskovg/typecode/server/src/internal/utils"
	"net/http"
	"github.com/sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/kletskovg/typecode/server/src/internal/github"
)

type APIServer struct{
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *APIServer) Start() error {
	if err := s.configureLogger(); err !=nil {
		return err
	}

	s.configureRouter()

	s.logger.Info("starting api server")
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)
	return nil
}

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/file", s.handleGetFile())
	s.router.HandleFunc("/", s.handleHello())
}

func (s *APIServer) handleHello() http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		go github.GetRandomRepository("JavaScript")
		w.Write([]byte(("Hello on Hello page")))
	}
}

func (s *APIServer) handleGetFile() http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		github.GetFile("go")
		
		w.Write([]byte(("Look to the console")))
	}
}
