package logic

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"product-srv/config"
	"product-srv/model"

	"product-srv/internal/svc"
	"product-srv/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCartLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCartLogic {
	return &AddCartLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddCartLogic) AddCart(in *product.AddCartReq) (*product.AddCartResp, error) {
	var user model.User
	if err := user.GetUserById(config.DB, in.Uid); err != nil {
		return nil, errors.New("用户查询失败")
	}
	if user.ID == 0 {
		return nil, errors.New("用户不存在")
	}

	//商品信息校验
	var Spu model.Spu
	if err := Spu.GetSpuBySpuId(config.DB, in.SpuId); err != nil {
		return nil, errors.New("商品查询失败")
	}
	if Spu.ID == 0 {
		return nil, errors.New("商品不存在")
	}
	var Sku model.Sku
	if err := Sku.GetSkuBSkuId(config.DB, in.SkuId); err != nil {
		return nil, errors.New("sku查询失败")
	}
	if Sku.ID == 0 {
		return nil, errors.New("sku不存在")
	}

	if int(in.Count) > Sku.Stock {
		return nil, errors.New("库存不足")
	}

	SpuMap := map[string]interface{}{
		"spu_id":   in.SpuId,
		"sku_id":   in.SkuId,
		"spu_name": Spu.SpuName,
		"sku_name": Sku.SkuName,
		"price":    Sku.Price,
		"stock":    in.Count,
	}

	CartInfo, _ := json.Marshal(SpuMap)

	key := fmt.Sprintf("cart:%v:%v", in.Uid, in.SpuId)
	exists := config.Rdb.Exists(config.Ctx, key)
	if exists.Val() == 0 {
		//不存在 创建
		err := config.Rdb.HSet(config.Ctx, key, "商品信息", string(CartInfo)).Err()
		if err != nil {
			return nil, errors.New("购物车添加失败")
		}

	} else {
		//存在 更新
		err := config.Rdb.HDel(config.Ctx, key, "商品信息").Err()
		if err != nil {
			return nil, errors.New("购物车更新失败-1")
		}
		err = config.Rdb.HSet(config.Ctx, key, "商品信息", string(CartInfo)).Err()
		if err != nil {
			return nil, errors.New("购物车更新失败-2")
		}
	}
	return &product.AddCartResp{}, nil
}
