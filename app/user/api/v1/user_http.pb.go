// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.4
// - protoc             v3.19.4
// source: v1/user.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationUserCompanyRegister = "/api.v1.User/CompanyRegister"

type UserHTTPServer interface {
	CompanyRegister(context.Context, *CompanyRegisterRequest) (*CompanyRegisterReply, error)
}

func RegisterUserHTTPServer(s *http.Server, srv UserHTTPServer) {
	r := s.Route("/")
	r.POST("/user/v1/company/register", _User_CompanyRegister0_HTTP_Handler(srv))
}

func _User_CompanyRegister0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CompanyRegisterRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserCompanyRegister)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CompanyRegister(ctx, req.(*CompanyRegisterRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CompanyRegisterReply)
		return ctx.Result(200, reply)
	}
}

type UserHTTPClient interface {
	CompanyRegister(ctx context.Context, req *CompanyRegisterRequest, opts ...http.CallOption) (rsp *CompanyRegisterReply, err error)
}

type UserHTTPClientImpl struct {
	cc *http.Client
}

func NewUserHTTPClient(client *http.Client) UserHTTPClient {
	return &UserHTTPClientImpl{client}
}

func (c *UserHTTPClientImpl) CompanyRegister(ctx context.Context, in *CompanyRegisterRequest, opts ...http.CallOption) (*CompanyRegisterReply, error) {
	var out CompanyRegisterReply
	pattern := "/user/v1/company/register"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserCompanyRegister))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
