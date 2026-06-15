package main

import (
	"flag"
	"fmt"
	"github.com/magic-lib/go-plat-utils/crypto"
)

func main() {
	flag.Usage = func() {
		fmt.Println("create_secret - 字符串加解密工具")
		fmt.Println()
		fmt.Println("用法:")
		fmt.Println("  create_secret -s <明文字符串> [-key <密钥>]")
		fmt.Println()
		fmt.Println("参数:")
		fmt.Println("  -s    原始明文字符串（必填）")
		fmt.Println("  -key  加密密钥（可选，默认为空）")
		fmt.Println()
		fmt.Println("示例:")
		fmt.Println("  create_secret -s \"mySecret\"")
		fmt.Println("  create_secret -s \"mySecret\" -key \"myKey\"")
	}

	originSecret := flag.String("s", "", "原始明文字符串")
	originKey := flag.String("key", "", "加密密钥（可选，默认为空）")
	flag.Parse()

	if *originSecret == "" {
		flag.Usage()
		return
	}

	encryptedStr, err := crypto.ConfigEncryptSecret(*originSecret, *originKey)
	if err != nil {
		fmt.Println("加密失败:", err)
		return
	}
	decryptedStr, err := crypto.ConfigDecryptSecret(encryptedStr, *originKey)
	if err != nil {
		fmt.Println("解密失败:", err)
		return
	}
	fmt.Println("原始字符串：", *originSecret)
	fmt.Println("加密字符串：", encryptedStr)
	fmt.Println("解密字符串：", decryptedStr)
}
