package productattributeservicelogic

import (
	"context"
	"zero-admin/rpc/model/pmsmodel"
	"zero-admin/rpc/pms/pmsclient"

	"zero-admin/rpc/pms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductAttributeAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductAttributeAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductAttributeAddLogic {
	return &ProductAttributeAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductAttributeAddLogic) ProductAttributeAdd(in *pmsclient.ProductAttributeAddReq) (*pmsclient.ProductAttributeAddResp, error) {
	_, err := l.svcCtx.PmsProductAttributeModel.Insert(l.ctx, &pmsmodel.PmsProductAttribute{
		ProductAttributeCategoryId: in.ProductAttributeCategoryId,
		Name:                       in.Name,
		SelectType:                 in.SelectType,
		InputType:                  in.InputType,
		InputList:                  in.InputList,
		Sort:                       in.Sort,
		FilterType:                 in.FilterType,
		SearchType:                 in.SearchType,
		RelatedStatus:              in.RelatedStatus,
		HandAddStatus:              in.HandAddStatus,
		Type:                       in.Type,
	})
	if err != nil {
		return nil, err
	}

	return &pmsclient.ProductAttributeAddResp{}, nil
}
