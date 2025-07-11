package service

import (
	"context"
	"fmt"
	v1 "kratos-demo/api/user/v1"
)

func (s *UserService) CompanyRegister(ctx context.Context, req *v1.CompanyRegisterRequest) (*v1.CompanyRegisterReply, error) {
	return s.companyBiz.Register(req)
	//return &v1.CompanyRegisterReply{}, nil
}

func (s *UserService) AccessToken(ctx context.Context, req *v1.AccessTokenRequest) (*v1.AccessTokenReply, error) {
	return s.companyBiz.CreateAccessToken(req)
}

func (s *UserService) VerifyAccessToken(ctx context.Context, req *v1.VerifyAccessTokenRequest) (*v1.VerifyAccessTokenReply, error) {
	if req.Token == "" || req.AppId == "" {
		return nil, fmt.Errorf("token or app_id is empty")
	}
	return s.companyBiz.VerifyAccessToken(req)
}
