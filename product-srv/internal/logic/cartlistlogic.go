package logic

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"product-srv/config"

	"product-srv/internal/svc"
	"product-srv/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type CartListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCartListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CartListLogic {
	return &CartListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CartListLogic) CartList(in *product.CartListReq) (*product.CartListResp, error) {
	var CartList []*product.CartInfo

	key := fmt.Sprintf("cart:%v:*", in.Uid)
	keys, _ := config.Rdb.Keys(config.Ctx, key).Result()
	if len(keys) == 0 {
		return nil, errors.New("当前用户暂无购物车信息")
	}
	for _, key := range keys {
		Info, _ := config.Rdb.HGet(config.Ctx, key, "商品信息").Result()
		var cartInfo product.CartInfo
		json.Unmarshal([]byte(Info), &cartInfo)
		fmt.Println(Info)
		CartList = append(CartList, &product.CartInfo{
			Price:   cartInfo.Price,
			SkuId:   cartInfo.SkuId,
			SkuName: cartInfo.SkuName,
			SpuId:   cartInfo.SpuId,
			SpuName: cartInfo.SpuName,
			Stock:   cartInfo.Stock,
		})
	}

	for i := 0; i < len(CartList); i++ {
		for j := i + 1; j < len(CartList); j++ {
			if CartList[i].Price > CartList[j].Price {
				value := CartList[i].Price
				CartList[i].Price = CartList[j].Price
				CartList[j].Price = value
			}
		}
	}

	return &product.CartListResp{
		List: CartList,
	}, nil
}
