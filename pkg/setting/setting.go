// setting 高性能选项集
// 时间 2019年3月9号
// 作者 申杨杨

package setting

// Redis 选项集
type Redis struct {
	Address  map[string]string // Redis 主机
	Password string            // Redis 密码
}

// 全局选项集
var (
	RedisSetting = &Redis{
		Address: map[string]string{
			"server1": ":6379",
		},
		Password: "",
	}
)
