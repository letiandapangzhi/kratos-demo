package service

import (
	"kratos-demo/app/user/internal/biz"

	v1 "kratos-demo/api/user/v1"
)

type UserService struct {
	v1.UnimplementedUserServer

	companyBiz *biz.CompanyUseCase
}

func NewUserService(cuc *biz.CompanyUseCase) *UserService {
	return &UserService{
		companyBiz: cuc,
	}
}
