package svc

import (
	"api/internal/config"
	"github.com/zeromicro/go-zero/zrpc"
	"product-srv/productclient"
)

type ServiceContext struct {
	Config     config.Config
	ProductSrv productclient.Product
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		ProductSrv: productclient.NewProduct(zrpc.MustNewClient(c.ProductSrv)),
	}
}
