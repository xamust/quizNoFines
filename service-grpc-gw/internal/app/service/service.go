package service

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/sirupsen/logrus"
	gateway "github.com/xamust/quizNoFines/service-grpc-gw/api"
	"google.golang.org/grpc"
	"net/http"
)

type Service struct {
	config *Config
	logger *logrus.Logger
}

// init new server
func New(config *Config) *Service {
	return &Service{
		config: config,
		logger: logrus.New(),
	}
}

// configure logrus...
func (s *Service) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}

func (s *Service) Start() error {

	//configure logger...
	if err := s.configureLogger(); err != nil {
		return err
	}

	//gRPC gateway
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gateway.RegisterClubsInfoHandlerFromEndpoint(ctx, mux, s.config.PortgRPC, opts)
	if err != nil {
		s.logger.Error(err)
		return err
	}
	s.logger.Infof("Starting gRPC dial on port: %v, starting listen on port: %v...", s.config.PortgRPC, s.config.Port)
	if err = http.ListenAndServe(s.config.Port, mux); err != nil {
		s.logger.Error(err)
		return err
	}
	return nil
}
