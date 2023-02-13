package redis

import (
	"context"
	"github.com/PICOF/simple-tiktok/pkg/config"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

var (
	Redis       *redis.Client
	redisConfig *viper.Viper
)

func init() {
	redisConfig = config.GetConfig("redis")
	Redis = redis.NewClient(&redis.Options{
		//连接信息
		Network:  redisConfig.GetString("network"), //网络类型，tcp or unix，默认tcp
		Addr:     redisConfig.GetString("address"),
		Password: redisConfig.GetString("password"), //密码
		DB:       redisConfig.GetInt("DBIndex"),     // redis数据库index

		//连接池容量及闲置连接数量
		PoolSize:     redisConfig.GetInt("poolSize"),     // 连接池最大socket连接数，默认为4倍CPU数， 4 * runtime.NumCPU
		MinIdleConns: redisConfig.GetInt("minIdleConns"), //在启动阶段创建指定数量的Idle连接，并长期维持idle状态的连接数不少于指定数量；。

		//钩子函数
		OnConnect: func(conn *redis.Conn) error { //仅当客户端执行命令时需要从连接池获取连接时，如果连接池需要新建连接时则会调用此钩子函数
			klog.CtxInfof(context.Background(), "conn=%v\n", conn)
			return nil
		},
	})
}
