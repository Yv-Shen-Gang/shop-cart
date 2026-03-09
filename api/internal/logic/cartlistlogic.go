package logic

import (
	"context"
	"product-srv/product"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CartlistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCartlistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CartlistLogic {
	return &CartlistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CartlistLogic) Cartlist(req *types.CartListReq) (resp *types.Response, err error) {
	cartList, err := l.svcCtx.ProductSrv.CartList(l.ctx, &product.CartListReq{
		Uid: req.Uid,
	})
	if err != nil {
		return &types.Response{
			Code: 400,
			Msg:  err.Error(),
			Data: nil,
		}, nil
	}

	return &types.Response{
		Code: 200,
		Msg:  "列表获取成功",
		Data: cartList,
	}, nil
}
