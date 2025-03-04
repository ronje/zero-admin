package cart

import (
	"context"
	"encoding/json"
	"zero-admin/rpc/oms/omsclient"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// CartUpdateQuantityLogic
/*
Author: LiuFeiHua
Date: 2023/12/6 15:15
*/
type CartUpdateQuantityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCartUpdateQuantityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CartUpdateQuantityLogic {
	return &CartUpdateQuantityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// CartUpdateQuantity 修改购物车中某个商品的数量
func (l *CartUpdateQuantityLogic) CartUpdateQuantity(req *types.CartItemUpdateQuantityReq) (resp *types.CartItemUpdateResp, err error) {
	memberId, _ := l.ctx.Value("memberId").(json.Number).Int64()
	_, _ = l.svcCtx.CartItemService.CartItemUpdateQuantity(l.ctx, &omsclient.CartItemUpdateReq{Id: req.Id, Quantity: req.Quantity, MemberId: memberId})

	return &types.CartItemUpdateResp{
		Code:    0,
		Message: "操作成功",
	}, nil
}
