package startupcfg

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"text/template"
)

// EncryptSecretFormat 组合解密和模板替换操作，还原密码配置
func EncryptSecretFormat(cfgTemplate string, encryptedMap map[string]Encrypted) (string, error) {
	if encryptedMap == nil || len(encryptedMap) == 0 {
		return cfgTemplate, nil
	}

	decodeSecretMap := make(map[string]string)
	for key, encryptedValue := range encryptedMap {
		decodedValue, err := encryptedValue.Get()
		if err != nil {
			return "", fmt.Errorf("failed to decode hex string for key %s: %w", key, err)
		}
		decodeSecretMap[key] = decodedValue
	}

	// 计算配置模板的 MD5 哈希值
	hash := md5.New()
	hash.Write([]byte(cfgTemplate))
	sum := hash.Sum(nil)

	var err error
	templateInstance := template.New(fmt.Sprintf("%x", sum))
	templateInstance = templateInstance.Option("missingkey=zero")
	templateInstance, err = templateInstance.Parse(cfgTemplate)
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %w", err)
	}
	var outputBuffer bytes.Buffer
	err = templateInstance.Execute(&outputBuffer, decodeSecretMap)
	if err != nil {
		return "", fmt.Errorf("failed to execute template: %w", err)
	}
	return outputBuffer.String(), nil
}

// EncryptSecretCreate 对密钥映射进行加密
func EncryptSecretCreate(secretMap map[string]string) (map[string]Encrypted, error) {
	encryptedMap := make(map[string]Encrypted)
	var retErr error
	for key, value := range secretMap {
		encryptedStr, err := Encrypted(value).Encode()
		if err != nil {
			return nil, fmt.Errorf("failed to encrypt value for key %s: %w", key, err)
		}
		encryptedMap[key] = encryptedStr
	}
	return encryptedMap, retErr
}
