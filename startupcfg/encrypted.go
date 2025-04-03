package startupcfg

import (
	"fmt"
	"github.com/magic-lib/go-plat-utils/crypto"
)

type (
	Encrypted string // Encrypted 加密串
)

var (
	encryptFunc = func(e string) (Encrypted, error) {
		return Encrypted(e), nil
	} // 默认将字符串转化为Encrypted类型
	hasSetEncryptHandler = false
	hasSetDecryptHandler = false // 是否设置过解密函数，一个程序里只能设置一次，不然所有以前的就都解不开了
	decryptFunc          = func(e Encrypted) (string, error) {
		return string(e), nil
	} // 默认的密码直接返回
)

// setEncryptHandler 设置加密函数
func setEncryptHandler(encryptF func(m string) (Encrypted, error)) {
	if hasSetEncryptHandler {
		return
	}
	if encryptF != nil {
		encryptFunc = encryptF
		hasSetEncryptHandler = true
	}
}
func setDecryptHandler(decryptF func(e Encrypted) (string, error)) {
	if hasSetDecryptHandler {
		return
	}

	if decryptF != nil {
		decryptFunc = decryptF
		hasSetDecryptHandler = true
	}
}

// SetDefaultEncryptedHandler 给默认的加解密方法设置加密key
func SetDefaultEncryptedHandler(key string) error {
	if hasSetDecryptHandler {
		return fmt.Errorf("decryptFunc has seted")
	}
	return SetEncryptAndDecryptHandler(func(m string) (Encrypted, error) {
		encryptedStr, err := crypto.ConfigEncryptSecret(m, key)
		if err != nil {
			return "", err
		}
		return Encrypted(encryptedStr), nil
	}, func(m Encrypted) (string, error) {
		str, err := crypto.ConfigDecryptSecret(string(m), key)
		if err != nil {
			return "", err
		}
		return str, nil
	})
}

// SetEncryptAndDecryptHandler 设置解密函数
func SetEncryptAndDecryptHandler(encryptF func(e string) (Encrypted, error),
	decryptF func(m Encrypted) (string, error)) error {
	if hasSetDecryptHandler {
		return fmt.Errorf("decryptFunc has seted")
	}
	setEncryptHandler(encryptF)
	setDecryptHandler(decryptF)
	return nil
}

// Get 获取解密串
func (e Encrypted) Get() (string, error) {
	if string(e) == "" {
		return "", nil
	}
	if decryptFunc != nil {
		return decryptFunc(e)
	}
	return string(e), fmt.Errorf("no set decryptFunc")
}

// Encode 加密
func (e Encrypted) Encode() (Encrypted, error) {
	if e == "" {
		return "", nil
	}
	if decryptFunc != nil {
		str, err := decryptFunc(e)
		if err == nil && str != string(e) {
			return e, nil
		}
	}

	if encryptFunc != nil {
		return encryptFunc(string(e))
	}
	return e, fmt.Errorf("no set encryptFunc")
}
