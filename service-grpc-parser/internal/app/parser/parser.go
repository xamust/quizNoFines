package parser

import "github.com/xamust/quizNoFines/service-grpc-parser/internal/app/model"

type Parser struct {
}

func CollectorInit() *Parser {
	return &Parser{}
}

func (p *Parser) Search(inn string) (*model.Data, error) {
	return nil, nil
}
