// Code generated by goctl. DO NOT EDIT.
// Source: ums.proto

package server

import (
	"context"

	"zero-admin/rpc/ums/internal/logic/memberstatisticsinfoservice"
	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/umsclient"
)

type MemberStatisticsInfoServiceServer struct {
	svcCtx *svc.ServiceContext
	umsclient.UnimplementedMemberStatisticsInfoServiceServer
}

func NewMemberStatisticsInfoServiceServer(svcCtx *svc.ServiceContext) *MemberStatisticsInfoServiceServer {
	return &MemberStatisticsInfoServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *MemberStatisticsInfoServiceServer) MemberStatisticsInfoAdd(ctx context.Context, in *umsclient.MemberStatisticsInfoAddReq) (*umsclient.MemberStatisticsInfoAddResp, error) {
	l := memberstatisticsinfoservicelogic.NewMemberStatisticsInfoAddLogic(ctx, s.svcCtx)
	return l.MemberStatisticsInfoAdd(in)
}

func (s *MemberStatisticsInfoServiceServer) MemberStatisticsInfoList(ctx context.Context, in *umsclient.MemberStatisticsInfoListReq) (*umsclient.MemberStatisticsInfoListResp, error) {
	l := memberstatisticsinfoservicelogic.NewMemberStatisticsInfoListLogic(ctx, s.svcCtx)
	return l.MemberStatisticsInfoList(in)
}

func (s *MemberStatisticsInfoServiceServer) MemberStatisticsInfoUpdate(ctx context.Context, in *umsclient.MemberStatisticsInfoUpdateReq) (*umsclient.MemberStatisticsInfoUpdateResp, error) {
	l := memberstatisticsinfoservicelogic.NewMemberStatisticsInfoUpdateLogic(ctx, s.svcCtx)
	return l.MemberStatisticsInfoUpdate(in)
}

func (s *MemberStatisticsInfoServiceServer) MemberStatisticsInfoDelete(ctx context.Context, in *umsclient.MemberStatisticsInfoDeleteReq) (*umsclient.MemberStatisticsInfoDeleteResp, error) {
	l := memberstatisticsinfoservicelogic.NewMemberStatisticsInfoDeleteLogic(ctx, s.svcCtx)
	return l.MemberStatisticsInfoDelete(in)
}
