package service

import (
	"github.com/sirupsen/logrus"
	sgRPC "github.com/xamust/quizNoFines/service-grpc-parser/api"
	"github.com/xamust/quizNoFines/service-grpc-parser/internal/app/parser"
	"google.golang.org/grpc"
	"net"
)

type Service struct {
	config *Config
	logger *logrus.Logger
	parse  *parser.Parser
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

	//gRPC server...
	listengRPS, err := net.Listen("tcp", s.config.PortgRPC)
	if err != nil {
		s.logger.Error(err)
		return err
	}

	serviceCollect := grpc.NewServer()
	sgRPC.RegisterClubsInfoServer(serviceCollect, &server{logger: s.logger, parser: s.parse})
	s.logger.Infof("Starting gRPC listener on port: %v ...", s.config.PortgRPC)
	if err = serviceCollect.Serve(listengRPS); err != nil {
		s.logger.Error("failed to serve: %v", err)
		return err
	}
	return nil
}
