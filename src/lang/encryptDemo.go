package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

func main() {
	data := encodeMsg("sdf3", "iscConfigService")
	// 9xgDQRImgwjos9dKnKae0Q==
	// java: bfWp5JjIf2TcFSteURVgqg==
	fmt.Println(data)

	act := decodeMsg(data, "iscConfigService")
	// sdf3
	fmt.Println(act)
}

func padding(src []byte, blocksize int) []byte {
	padnum := blocksize - len(src)%blocksize
	pad := bytes.Repeat([]byte{byte(padnum)}, padnum)
	return append(src, pad...)
}

func unpadding(src []byte) []byte {
	n := len(src)
	unpadnum := int(src[n-1])
	return src[:n-unpadnum]
}

func EncryptAES2(src []byte, key []byte) []byte {
	block, _ := aes.NewCipher(key)
	src = padding(src, block.BlockSize())
	blockmode := cipher.NewCBCEncrypter(block, key)
	blockmode.CryptBlocks(src, src)
	return src
}

func DecryptAES2(src []byte, key []byte) []byte {
	block, _ := aes.NewCipher(key)
	blockmode := cipher.NewCBCDecrypter(block, key)
	blockmode.CryptBlocks(src, src)
	src = unpadding(src)
	return src
}

func decodeMsg(result, key string) string {
	middle, err := Base64Decode(result)
	if err != nil {
		fmt.Println("err1: ", err.Error())
		return ""
	}
	re := DecryptAES2(middle, []byte(key))
	return string(re)
}

func encodeMsg(content, key string) string {
	re := EncryptAES2([]byte(content), []byte(key))
	return Base64Encode(re)
}

func Base64Encode(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}

func Base64Decode(dst string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(dst)
}
