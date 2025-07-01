package biz

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	v1 "kratos-demo/api/user/v1"
	"kratos-demo/app/user/internal/conf"
	"kratos-demo/pkg/util"
	"strconv"
	"strings"
)

type Company struct {
	ID        int64  `json:"id"`
	AppID     string `json:"appId,omitempty"`
	AppSecret string `json:"appSecret"`
}

// CompanyRepo company repository
type CompanyRepo interface {
	Save(*v1.CompanyRegisterRequest) (*Company, error)
	GetInfo(where *Company) (*Company, error)

	SaveWhiteHash(*v1.AccessTokenRequest) error
}

type CompanyUseCase struct {
	auth *conf.Auth
	repo CompanyRepo
	log  *log.Helper
}

func NewCompanyUseCase(repo CompanyRepo, auth *conf.Auth, logger log.Logger) *CompanyUseCase {
	return &CompanyUseCase{repo: repo, auth: auth, log: log.NewHelper(logger)}
}

// Register 注册
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

// CreateAccessToken 获取访问令牌
func (uc *CompanyUseCase) CreateAccessToken(req *v1.AccessTokenRequest) (*v1.AccessTokenReply, error) {
	// todo 验证时间戳有效性
	// 查询用户
	info, err := uc.repo.GetInfo(&Company{AppID: req.AppId})
	if err != nil {
		return nil, err
	} else if info.ID == 0 {
		return nil, errors.New("主体不存在")
	}
	mySign := util.SHA3256Hash(fmt.Sprintf("%s%s%s", req.AppId, info.AppSecret, req.Timestamp))
	if mySign != req.Sign {
		// 验证hash appId+appSecret+时间戳
		return nil, errors.New("签名错误")
	}
	// 生成token
	token, err := uc.accessToken(info)
	if err != nil {
		return nil, err
	}
	// redis白名单
	err = uc.repo.SaveWhiteHash(req)
	if err != nil {
		return nil, err
	}
	return &v1.AccessTokenReply{AccessToken: token}, nil
}

func (uc *CompanyUseCase) VerifyAccessToken(req *v1.VerifyAccessTokenRequest) (*v1.VerifyAccessTokenReply, error) {
	// 验证token
	_, err := uc.checkAccessToken(req.Token)
	if err != nil {
		return nil, err
	}
	// todo 有效性
	return &v1.VerifyAccessTokenReply{
		Valid: true,
	}, nil
}

// accessToken 整合加密生成访问令牌
func (uc *CompanyUseCase) accessToken(company *Company) (string, error) {
	key, _ := hex.DecodeString(uc.auth.Aes256Key)
	// aes256加密
	ciphertext, nonce, err := util.Aes256GCMEncrypt(key, []byte(company.AppSecret))
	if err != nil {
		return "", err
	}
	// 拼接混淆 版本 ID 盐 密文
	return fmt.Sprintf("%d%dz%x%xz%x", 1, company.ID+int64(nonce[0]), ^nonce[0], nonce[1:], ciphertext), nil
}

// checkAccessToken 整合解密访问令牌
func (uc *CompanyUseCase) checkAccessToken(token string) (*Company, error) {
	tokenSlice := strings.Split(token, "z")
	// 混淆版本
	//version := tokenSlice[0][0:1]
	id, err := strconv.ParseInt(tokenSlice[0][1:], 10, 64)
	if err != nil {
		return nil, err
	}
	// 盐
	nonce, err := hex.DecodeString(tokenSlice[1])
	if err != nil {
		return nil, err
	}
	nonce[0] = ^nonce[0]
	// 密钥
	key, _ := hex.DecodeString(uc.auth.Aes256Key)
	// 密文
	ciphertext, _ := hex.DecodeString(tokenSlice[2])
	// aes256解密
	plaintext, err := util.Aes256GCMDecrypt(key, nonce, ciphertext)

	company := &Company{
		ID:        id - int64(nonce[0]),
		AppSecret: string(plaintext),
	}
	return company, nil
}
