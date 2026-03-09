package config

import (
	"context"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var Config *AddConfig
var DB *gorm.DB
var Rdb *redis.Client
var Ctx = context.Background()
