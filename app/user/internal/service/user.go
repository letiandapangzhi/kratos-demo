package service

import (
	"context"
	"fmt"
	"kratos-demo/app/user/internal/biz"

	pb "kratos-demo/api/user/v1"
)

type UserService struct {
	pb.UnimplementedUserServer

	userUc *biz.UserUsecase
}

func NewUserService(user *biz.UserUsecase) *UserService {
	return &UserService{
		userUc: user,
	}
}

func (s *UserService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterReply, error) {
	userModel := &biz.User{}
	err := s.userUc.Register(ctx, userModel)
	if err != nil {
		return nil, err
	}
	return &pb.RegisterReply{
		Message: fmt.Sprintf("%d", userModel.ID),
	}, nil
}
