package startupcfg_test

import (
	"fmt"
	"github.com/magic-lib/go-plat-startupcfg/startupcfg"
	"testing"
)

func TestEncodeSecretMap(t *testing.T) {
	_ = startupcfg.SetDefaultEncryptedHandler("aa")
	newMap, err := startupcfg.EncryptSecretCreate(map[string]string{
		"aaa": "kkk",
	})
	fmt.Println(newMap, err)
}
func TestFormatWithSecretMap(t *testing.T) {
	_ = startupcfg.SetDefaultEncryptedHandler("aa")
	newMap, err := startupcfg.DecryptSecretFormat("aaaaa{{.aaa}}", map[string]startupcfg.Encrypted{
		"aaa": "635d662e90ba51615e811efed2f8be98",
	})
	fmt.Println(newMap, err)
}
