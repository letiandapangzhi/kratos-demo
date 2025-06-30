package data

import (
	"encoding/hex"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	v1 "kratos-demo/api/user/v1"
	"kratos-demo/app/user/internal/biz"
	"kratos-demo/app/user/internal/data/gen/model"
	"kratos-demo/pkg/util"
	"time"
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
	secret, err := util.RandomByte(16)
	if err != nil {
		return nil, err
	}
	password, err := util.BcryptHash(util.SHA3256Hash(req.Phone[5:]))
	if err != nil {
		return nil, err
	}
	insertData := &model.Company{
		Name:      req.Name,
		Phone:     req.Phone,
		Password:  password,
		AppID:     fmt.Sprintf("xt%d", time.Now().UnixMilli()),
		AppSecret: hex.EncodeToString(secret), // 16字节转16进制字符串长度是32
	}
	err = r.data.model.Company.Create(insertData)
	if err != nil {
		return nil, err
	}
	return &biz.Company{
		AppID:     insertData.AppID,
		AppSecret: insertData.AppSecret,
	}, nil
}

func (r *companyRepo) GetInfo(*biz.Company) (*biz.Company, error) {
	return &biz.Company{}, nil
}

func (r *companyRepo) SaveWhiteHash(*v1.AccessTokenRequest) error {
	return nil
}
