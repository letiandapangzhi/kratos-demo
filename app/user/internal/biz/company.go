package biz

import (
	v1 "kratos-demo/api/user/v1"

	"github.com/go-kratos/kratos/v2/log"
)

type Company struct {
	AppID     string
	AppSecret string
}

// CompanyRepo company repository
type CompanyRepo interface {
	Save(reply *v1.CompanyRegisterRequest) (*Company, error)
}

type CompanyUseCase struct {
	repo CompanyRepo
	log  *log.Helper
}

func NewCompanyUseCase(repo CompanyRepo, logger log.Logger) *CompanyUseCase {
	return &CompanyUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *CompanyUseCase) Register(req *v1.CompanyRegisterRequest) (*v1.CompanyRegisterReply, error) {
	company, err := uc.repo.Save(req)
	if err != nil {
		return nil, err
	}
	return &v1.CompanyRegisterReply{
		AppId:     company.AppID,
		AppSecret: company.AppSecret,
	}, nil
}
