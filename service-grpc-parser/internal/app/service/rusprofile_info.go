package service

import (
	"context"
	"github.com/xamust/quizNoFines/service-grpc-parser/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	Service *Service
}

func (s *server) GetCompany(ctx context.Context, company *api.INNCompany) (*api.CompanyInfo, error) {
	_, err := s.Service.parse.Search(company.INN)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error while search company by INN %s, err =  %v", company.INN, err)
	}
	return &api.CompanyInfo{
		INN: "111",
		//INN:             search.INN,
		KPP: "222",
		//KPP:             search.KPP,
		CompanyName: "test companyName",
		//CompanyName:     search.CompanyName,
		FullNameManager: "Ivan Ivanov",
		//FullNameManager: search.FullNameManager,
	}, status.New(codes.OK, "").Err()
}
