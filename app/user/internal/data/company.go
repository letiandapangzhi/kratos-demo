package data

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gen"
	"gorm.io/gorm"
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

func (r *companyRepo) GetInfo(req *biz.Company) (*biz.Company, error) {
	company := r.data.model.Company
	where := make([]gen.Condition, 0)
	if req.ID != 0 {
		where = append(where, company.ID.Eq(int32(req.ID)))
	}
	if req.AppID != "" {
		where = append(where, company.AppID.Eq(req.AppID))
	}
	if req.AppSecret != "" {
		where = append(where, company.AppSecret.Eq(req.AppSecret))
	}
	info, err := company.Where(where...).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &biz.Company{}, nil
		}
		return nil, err
	}
	return &biz.Company{
		ID:        int64(info.ID),
		AppID:     info.AppID,
		AppSecret: info.AppSecret,
	}, nil
}

func (r *companyRepo) SaveWhiteHash(*v1.AccessTokenRequest) error {
	return nil
}
