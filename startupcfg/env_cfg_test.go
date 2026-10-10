package startupcfg_test

import (
	"os"
	"testing"

	"github.com/magic-lib/go-plat-startupcfg/startupcfg"
)

const envTestPrefix = "TEST_ENV_CFG_"

func prepareTestEnv(t *testing.T) map[string]string {
	t.Helper()
	testEnv := map[string]string{
		envTestPrefix + "NAME": "gdp",
		envTestPrefix + "PORT": "8080",
		envTestPrefix + "URL":  "http://www.baidu.com",
	}
	if err := startupcfg.SetMapToEnv(testEnv); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		for k := range testEnv {
			_ = os.Unsetenv(k)
		}
	})
	return testEnv
}

func TestSetMapToEnv(t *testing.T) {
	testEnv := map[string]string{
		envTestPrefix + "NAME": "gdp",
		envTestPrefix + "PORT": "8080",
	}
	if err := startupcfg.SetMapToEnv(testEnv); err != nil {
		t.Error(err)
		return
	}
	t.Cleanup(func() {
		for k := range testEnv {
			_ = os.Unsetenv(k)
		}
	})

	for k, v := range testEnv {
		if got := os.Getenv(k); got != v {
			t.Errorf("SetMapToEnv() key %s = %s, want %s", k, got, v)
		}
	}
}

func TestEnvToMapByPrefix(t *testing.T) {
	testEnv := prepareTestEnv(t)

	m := startupcfg.EnvToMapByPrefix("")
	for k, v := range testEnv {
		got, ok := m[k]
		if !ok {
			t.Errorf("EnvToMapByPrefix(%s) missing key %s", envTestPrefix, k)
			continue
		}
		if got != v {
			t.Errorf("EnvToMapByPrefix(%s) key %s = %s, want %s", envTestPrefix, k, got, v)
		}
	}
	if len(m) < len(testEnv) {
		t.Errorf("EnvToMapByPrefix(%s) got %d items, want at least %d", envTestPrefix, len(m), len(testEnv))
	}

	// 不存在的 prefix 返回空 map
	empty := startupcfg.EnvToMapByPrefix("TEST_ENV_CFG_NOT_EXIST_PREFIX_")
	if len(empty) != 0 {
		t.Errorf("EnvToMapByPrefix() with unknown prefix got %v, want empty", empty)
	}
}

func TestEnvToMap(t *testing.T) {
	testEnv := prepareTestEnv(t)

	m := startupcfg.EnvToMap()
	for k, v := range testEnv {
		got, ok := m[k]
		if !ok {
			t.Errorf("EnvToMap() missing key %s", k)
			continue
		}
		if got != v {
			t.Errorf("EnvToMap() key %s = %s, want %s", k, got, v)
		}
	}

	// 空 prefix 等价于 EnvToMap
	if len(m) != len(startupcfg.EnvToMapByPrefix("")) {
		t.Errorf("EnvToMap() got %d items, EnvToMapByPrefix(\"\") got %d items", len(m), len(startupcfg.EnvToMapByPrefix("")))
	}
}
