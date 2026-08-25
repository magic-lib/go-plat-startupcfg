package startupcfg

import "fmt"

// ServiceAPI 服务接口抽象
type ServiceAPI interface {
	// DomainName 接口域名
	DomainName() string
	// Url 接口Url
	Url(apiName string) string
	// AuthData 接口其他数据（鉴权数据等）
	AuthData(key string) (string, error)
	// ConfigData 获取服务所需的其它自定义参数
	ConfigData(key string) (any, bool)
	ConfigAll() map[string]any
}

// ServiceApiConfig 服务接口
type ServiceApiConfig struct {
	Domain string               `json:"domain" yaml:"domain"`
	Auth   map[string]Encrypted `json:"auth" yaml:"auth"`
	Urls   map[string]string    `json:"urls" yaml:"urls"`
	// Params 服务所需的其它自定义参数
	Configs map[string]any `json:"configs" yaml:"configs"`
}

// DomainName 接口域名
func (c *ServiceApiConfig) DomainName() string {
	return c.Domain
}

// Url 接口Url
func (c *ServiceApiConfig) Url(name string) string {
	if c.Urls != nil {
		if url, ok := c.Urls[name]; ok {
			return url
		}
	}
	return ""
}

// AuthData 接口其他数据（鉴权数据等）
func (c *ServiceApiConfig) AuthData(key string) (string, error) {
	if c == nil {
		return "", fmt.Errorf("auth data %s empty", key)
	}
	if c.Auth != nil {
		if valueEncrypt, ok := c.Auth[key]; ok {
			return valueEncrypt.Get()
		}
	}
	return "", nil
}

// ConfigData 获取服务所需的其它自定义参数
func (c *ServiceApiConfig) ConfigData(key string) (any, bool) {
	if c == nil || c.Configs == nil {
		return "", false
	}
	value, ok := c.Configs[key]
	return value, ok
}

// ConfigAll 获取服务所需的其它自定义参数
// 返回副本，避免调用方修改 map 覆盖内部配置
func (c *ServiceApiConfig) ConfigAll() map[string]any {
	if c == nil || c.Configs == nil {
		return map[string]any{}
	}
	result := make(map[string]any, len(c.Configs))
	for k, v := range c.Configs {
		result[k] = v
	}
	return result
}
