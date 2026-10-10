package startupcfg

import (
	"os"
	"strings"
)

func EnvToMap() map[string]string {
	return EnvToMapByPrefix("")
}
func EnvToMapByPrefix(prefix string) map[string]string {
	m := make(map[string]string)
	for _, s := range os.Environ() {
		if k, v, ok := strings.Cut(s, "="); ok {
			if prefix == "" {
				m[k] = v
				continue
			}
			if !strings.HasPrefix(k, prefix) {
				continue
			}
			m[k] = v
		}
	}
	return m
}

// SetMapToEnv 将map写入进程环境变量
func SetMapToEnv(m map[string]string) error {
	for k, v := range m {
		if err := os.Setenv(k, v); err != nil {
			return err
		}
	}
	return nil
}
