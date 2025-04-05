package startupcfg

type (
	DriverType  string // DriverType 连接类型
	ExtendField string // ExtendField 扩展字段名
)

const (
	defaultMysqlCharset = "utf8"
)

var (
	DriverMysql DriverType = "mysql"
	DriverRedis DriverType = "redis"

	extendMysqlCharset     ExtendField = "charset"
	extendRedisTLS         ExtendField = "tls"
	extendRedisType        ExtendField = "type"
	extendRedisPingTimeout ExtendField = "ping_timeout"

	_, _ Database = new(MysqlConfig), new(RedisConfig)
)
