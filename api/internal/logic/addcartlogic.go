package logic

import (
	"context"
	"product-srv/product"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCartLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCartLogic {
	return &AddCartLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddCartLogic) AddCart(req *types.AddCartReq) (resp *types.Response, err error) {
	_, err = l.svcCtx.ProductSrv.AddCart(l.ctx, &product.AddCartReq{
		SpuId: req.SpuId,
		SkuId: req.SkuId,
		Uid:   req.Uid,
		Count: req.Count,
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
		Msg:  "购物车添加成功",
		Data: nil,
	}, nil
}
