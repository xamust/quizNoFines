package parser

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/sirupsen/logrus"
	"github.com/xamust/quizNoFines/service-grpc-parser/internal/app/model"
)

const hrefPref = "https://www.rusprofile.ru/search?query=%s&search_inactive=0"

type Parser struct {
	logger *logrus.Logger
	c      *colly.Collector
}

func CollectorInit(logger *logrus.Logger) *Parser {
	return &Parser{
		c:      colly.NewCollector(),
		logger: logger,
	}
}

func (p *Parser) Search(inn string) (*model.Data, error) {
	companyData := &model.Data{}
	p.c.OnHTML("#clip_inn", func(e *colly.HTMLElement) {
		companyData.INN = e.Text
	})
	p.c.OnHTML("#clip_kpp", func(e *colly.HTMLElement) {
		companyData.KPP = e.Text
	})
	p.c.OnHTML(".company-name", func(e *colly.HTMLElement) {
		companyData.CompanyName = e.Text
	})
	p.c.OnHTML(".link-arrow.gtm_main_fl", func(e *colly.HTMLElement) {
		companyData.FullNameManager = e.Text
	})
	if err := p.c.Visit(fmt.Sprintf(hrefPref, inn)); err != nil {
		return nil, fmt.Errorf("error with parsing %s, error:%v", fmt.Sprintf(hrefPref, inn), err)
	}
	p.logger.Infof("data by inn: %v = %#v", inn, *companyData)

	return companyData, nil
}
