// Code generated by goctl. DO NOT EDIT.
// Source: sys.proto

package server

import (
	"context"

	"zero-admin/rpc/sys/internal/logic/loginlogservice"
	"zero-admin/rpc/sys/internal/svc"
	"zero-admin/rpc/sys/sysclient"
)

type LoginLogServiceServer struct {
	svcCtx *svc.ServiceContext
	sysclient.UnimplementedLoginLogServiceServer
}

func NewLoginLogServiceServer(svcCtx *svc.ServiceContext) *LoginLogServiceServer {
	return &LoginLogServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *LoginLogServiceServer) LoginLogAdd(ctx context.Context, in *sysclient.LoginLogAddReq) (*sysclient.LoginLogAddResp, error) {
	l := loginlogservicelogic.NewLoginLogAddLogic(ctx, s.svcCtx)
	return l.LoginLogAdd(in)
}

func (s *LoginLogServiceServer) LoginLogList(ctx context.Context, in *sysclient.LoginLogListReq) (*sysclient.LoginLogListResp, error) {
	l := loginlogservicelogic.NewLoginLogListLogic(ctx, s.svcCtx)
	return l.LoginLogList(in)
}

func (s *LoginLogServiceServer) LoginLogDelete(ctx context.Context, in *sysclient.LoginLogDeleteReq) (*sysclient.LoginLogDeleteResp, error) {
	l := loginlogservicelogic.NewLoginLogDeleteLogic(ctx, s.svcCtx)
	return l.LoginLogDelete(in)
}
