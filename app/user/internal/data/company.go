package data

import (
	v1 "kratos-demo/api/user/v1"
	"kratos-demo/app/user/internal/biz"
	"kratos-demo/app/user/internal/data/gen/model"

	"github.com/go-kratos/kratos/v2/log"
)

type companyRepo struct {
	data *Data
	log  *log.Helper
}

func NewCompanyRepo(data *Data, logger log.Logger) biz.CompanyRepo {
	return &companyRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *companyRepo) Save(req *v1.CompanyRegisterRequest) (*biz.Company, error) {
	insertData := &model.TestCompany{
		Name:  req.Name,
		Phone: req.Phone,
	}
	err := r.data.model.TestCompany.Create()
	if err != nil {
		return nil, err
	}
	return &biz.Company{
		AppID:     insertData.AppID,
		AppSecret: insertData.AppSecret,
	}, nil
}
