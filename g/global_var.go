package g

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"sync"
)

var Config *viper.Viper
var Redis *redis.Client
var Logger *zap.Logger
var Orm *gorm.DB                            // GORM
var RateLimits *sync.Map 					// 访问的IP map
