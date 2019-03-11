// redis 高性能工具集
// 时间 2019年3月9号
// 作者 申杨杨

package gredis

import (
	"github.com/go-redis/cache"
	"github.com/go-redis/redis"
	"gohighfaner/pkg/json"
	"gohighfaner/pkg/setting"
	"time"
)

// cache
var codec *cache.Codec

// Setup 函数
func Setup() {
	// 创建 redis 客户端
	ring := redis.NewRing(&redis.RingOptions{
		Addrs:    setting.RedisSetting.Address,
		Password: setting.RedisSetting.Password,
	})
	// 创建 cache 工具
	codec = &cache.Codec{
		Redis: ring,
		Marshal: func(v interface{}) ([]byte, error) {
			return json.Marshal(v)
		},
		Unmarshal: func(b []byte, v interface{}) error {
			return json.Unmarshal(b, v)
		},
	}
}

// Set ...
func Set(key string, value interface{}, expire time.Duration) error {
	return codec.Set(&cache.Item{
		Key:        key,
		Object:     value,
		Expiration: expire,
	})
}

// Get ...
func Get(key string, value interface{}) error {
	return codec.Get(key, value)
}
