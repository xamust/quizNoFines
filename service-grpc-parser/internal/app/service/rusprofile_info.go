package service

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/xamust/quizNoFines/service-grpc-parser/api"
	"github.com/xamust/quizNoFines/service-grpc-parser/internal/app/parser"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	parser *parser.Parser
	logger *logrus.Logger
}

func (s *server) GetCompany(ctx context.Context, company *api.INNCompany) (*api.CompanyInfo, error) {
	p := parser.CollectorInit(s.logger)
	search, err := p.Search(company.INN)
	if err != nil {
		s.logger.Errorf("error while search company by INN %s, err =  %v", company.INN, err)
		return nil, status.Errorf(codes.Internal, "error while search company by INN %s, err =  %v", company.INN, err)
	}
	return &api.CompanyInfo{
		INN:             search.INN,
		KPP:             search.KPP,
		CompanyName:     search.CompanyName,
		FullNameManager: search.FullNameManager,
	}, status.New(codes.OK, "").Err()
}
