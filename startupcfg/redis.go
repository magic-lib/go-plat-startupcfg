package startupcfg

import (
	"fmt"
	"log"
	"net/url"
)

var _ Database = (*RedisConfig)(nil)

// RedisConfig redis配置
type RedisConfig struct {
	Protocol        string    `json:"protocol" yaml:"protocol"`
	PasswordEncoded Encrypted `json:"pwEncoded" yaml:"pwEncoded"`
	Address         string    `json:"address" yaml:"address"`
	Database        int64     `json:"database" yaml:"database"`
	Username        string    `json:"username" yaml:"username"`
	UseTLS          bool      `json:"useTLS" yaml:"useTLS"`
}

// ProtocolName 连接协议
func (c *RedisConfig) ProtocolName() string {
	if c.Protocol == "" {
		c.Protocol = "tcp"
	}
	return c.Protocol
}

// DriverName 驱动名称
func (c *RedisConfig) DriverName() string {
	return string(DriverRedis)
}

// DatasourceName 连接数据库时的datasourceName参数
func (c *RedisConfig) DatasourceName() string {
	return fmt.Sprintf("%s://%s:%s@%s/%d",
		c.DriverName(),
		c.User(),
		url.QueryEscape(c.Password()),
		c.ServerAddress(),
		c.DatabaseName())
}

// ServerAddress redis服务器地址
func (c *RedisConfig) ServerAddress() string {
	return c.Address
}

// Password redis数据库用户密码
func (c *RedisConfig) Password() string {
	pass, err := c.PasswordEncoded.Get()
	if err != nil {
		log.Println("redis password decode error:", err)
	}
	return pass
}

// DatabaseName redis数据库名称
func (c *RedisConfig) DatabaseName() interface{} {
	return c.Database
}

// User redis数据库用户
func (c *RedisConfig) User() string {
	return c.Username
}

// Extend 扩展字段
func (c *RedisConfig) Extend(name ExtendField) (interface{}, bool) {
	switch name {
	case extendRedisUseTLS:
		return c.UseTLS, true
	}
	return nil, false
}
