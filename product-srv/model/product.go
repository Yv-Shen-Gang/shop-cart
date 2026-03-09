package model

import "gorm.io/gorm"

type Spu struct {
	gorm.Model
	SpuId   string `gorm:"type:varchar(30);comment:商品编号" json:"spuId"`
	SpuName string `gorm:"type:varchar(30);comment:商品名称" json:"spuName"`
	Info    string `gorm:"type:varchar(30);comment:商品简介" json:"info"`
}

func (s *Spu) GetSpuBySpuId(db *gorm.DB, id string) error {
	return db.Debug().Where("spu_id = ?", id).Limit(1).Find(s).Error
}

type Sku struct {
	gorm.Model
	SpuId   string  `gorm:"type:varchar(30);comment:商品编号" json:"spuId"`
	SkuId   string  `gorm:"type:varchar(30);comment:Sku编号" json:"skuId"`
	SkuName string  `gorm:"type:varchar(30);comment:Sku名称" json:"skuName"`
	Price   float64 `gorm:"type:decimal(10,2);comment:价格" json:"price"`
	Stock   int     `gorm:"type:int;comment:库存" json:"stock"`
}

func (s *Sku) GetSkuBSkuId(db *gorm.DB, id string) error {
	return db.Debug().Where("sku_id = ?", id).Limit(1).Find(&s).Error
}

type Specs struct {
	gorm.Model
	SkuId string `gorm:"type:varchar(30);comment:Sku编号" json:"skuId"`
	Key   string `gorm:"type:varchar(30);comment:规格键" json:"key"`
	Value string `gorm:"type:varchar(30);comment:规格值" json:"value"`
}
