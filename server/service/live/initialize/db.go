package initialize

import (
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jizizr/goligoli/server/common/consts"
	"github.com/jizizr/goligoli/server/service/live/config"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

// InitDB to init database
func InitDB() *gorm.DB {
	c := config.GlobalServerConfig.MysqlInfo
	dsn := fmt.Sprintf(consts.MySqlDSN, c.Username, c.Password, c.Host, c.Port, c.Name)
	// global mode
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		klog.Fatalf("init gorm failed: %s", err.Error())
	}
	return db
}

func InitRedis() *redis.Client {
	c := config.GlobalServerConfig.RedisInfo
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", c.Host, c.Port),
		Password: c.Password,
		DB:       c.DB,
	})
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	return rdb
}
