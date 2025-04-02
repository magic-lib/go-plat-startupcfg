package startupcfg

import (
	"fmt"
	"log"
	"net/url"
)

var _ Database = (*MysqlConfig)(nil)

// MysqlConfig mysql配置
type MysqlConfig struct {
	Protocol        string    `json:"protocol" yaml:"protocol"`
	UserName        string    `json:"username" yaml:"username"`
	PasswordEncoded Encrypted `json:"pwEncoded" yaml:"pwEncoded"`
	Address         string    `json:"address" yaml:"address"`
	Database        string    `json:"database" yaml:"database"`
	Charset         string    `json:"charset" yaml:"charset"`
}

// ProtocolName 连接协议
func (c *MysqlConfig) ProtocolName() string {
	return c.Protocol
}

// DriverName 使用sql.Open连接数据库时的driverName参数
func (c *MysqlConfig) DriverName() string {
	return string(DriverMysql)
}

// DatasourceName 使用sql.Open连接数据库时的datasourceName参数
func (c *MysqlConfig) DatasourceName() string {
	if c.Charset == "" {
		c.Charset = _MYSQL_CHARSET
	}
	//parseTime 是会自动解析成go的time.Time，否则为字符串
	//loc 时区设置：loc=Asia/Shanghai 则用东八区解析时间
	parseTime := "true"
	loc := "Local"

	if c.Protocol == "" {
		c.Protocol = "tcp"
	}

	return fmt.Sprintf("%s:%s@%s(%s)/%s?charset=%s&parseTime=%s&loc=%s",
		c.UserName,
		url.QueryEscape(c.Password()),
		c.Protocol,
		c.Address,
		c.Database,
		c.Charset,
		parseTime,
		loc)
}

// ServerAddress mysql服务器地址
func (c *MysqlConfig) ServerAddress() string {
	return c.Address
}

// User mysql数据库用户
func (c *MysqlConfig) User() string {
	return c.UserName
}

// Password mysql数据库用户密码
func (c *MysqlConfig) Password() string {
	pass, err := c.PasswordEncoded.Get()
	if err != nil {
		log.Println("mysql password decode error:", err)
	}
	return pass
}

// DatabaseName mysql数据库名称
func (c *MysqlConfig) DatabaseName() interface{} {
	return c.Database
}

// Extend 扩展字段
func (c *MysqlConfig) Extend(name ExtendField) (interface{}, bool) {
	if name == extendMysqlCharset {
		return c.Charset, true
	}
	return nil, false
}
