package service

import (
	"context"
	v1 "kratos-demo/app/user/api/v1"
)

func (s *UserService) CompanyRegister(ctx context.Context, req *v1.CompanyRegisterRequest) (*v1.CompanyRegisterReply, error) {
	return s.companyBiz.Register(req)
	//return &v1.CompanyRegisterReply{}, nil
}
