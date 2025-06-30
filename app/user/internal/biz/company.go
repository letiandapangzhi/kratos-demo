package biz

import (
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	v1 "kratos-demo/api/user/v1"
	"kratos-demo/pkg/util"
)

type Company struct {
	ID        int64
	AppID     string
	AppSecret string
}

// CompanyRepo company repository
type CompanyRepo interface {
	Save(*v1.CompanyRegisterRequest) (*Company, error)
	GetInfo(*Company) (*Company, error)

	SaveWhiteHash(*v1.AccessTokenRequest) error
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

func (uc *CompanyUseCase) CreateAccessToken(req *v1.AccessTokenRequest) (*v1.AccessTokenReply, error) {
	// 查询用户
	info, err := uc.repo.GetInfo(&Company{AppID: req.AppId})
	if err != nil {
		return nil, err
	} else if info.ID == 0 {
		return nil, errors.New("主体不存在")
	} else if req.Sign != util.SHA3256Hash(fmt.Sprintf("%s%s%s", req.AppId, info.AppSecret, req.Timestamp)) {
		// 验证hash appId+appSecret+时间戳
		return nil, errors.New("签名错误")
	}
	// 生成token
	token := uc.accessToken(info)
	// redis白名单
	err = uc.repo.SaveWhiteHash(req)
	if err != nil {
		return nil, err
	}
	return &v1.AccessTokenReply{AccessToken: token}, nil
}

// accessToken 整合加密生成访问令牌
func (uc *CompanyUseCase) accessToken(company *Company) string {
	// 配置读取密钥
	// aes256加密
	util.Aes256GCMEncrypt([]byte("key"), []byte("plaintext"))
	// 拼接打乱替换
	return ""
}
