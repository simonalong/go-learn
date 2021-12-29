package main

import (
	"fmt"
	"os/exec"
)

//
//import (
//	"bytes"
//	"crypto/aes"
//	"crypto/cipher"
//	"crypto/md5"
//	"encoding/base64"
//	"encoding/hex"
//	"fmt"
//)
//
//func generateMd5Key(str string) []byte {
//	checksum := md5.Sum([]byte(str))
//	return checksum[0:]
//}
//
//func aesEncryptECB(plaintext string, key string) string {
//	origData := []byte(plaintext)
//	checksum := generateMd5Key(key)
//	cipher, _ := aes.NewCipher(checksum)
//	length := (len(origData) + aes.BlockSize) / aes.BlockSize
//	plain := make([]byte, length*aes.BlockSize)
//	copy(plain, origData)
//	pad := byte(len(plain) - len(origData))
//	for i := len(origData); i < len(plain); i++ {
//		plain[i] = pad
//	}
//	encrypted := make([]byte, len(plain))
//	for bs, be := 0, cipher.BlockSize(); bs <= len(origData); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
//		cipher.Encrypt(encrypted[bs:be], plain[bs:be])
//	}
//
//	return Base64Encode(encrypted)
//}
//func aesDecryptECB(ciphertext string, key string) (string, error) {
//	checksum := generateMd5Key(key)
//	cipher, _ := aes.NewCipher(checksum)
//	encrypted, err := hex.DecodeString(ciphertext)
//	if err != nil {
//		return "", err
//	}
//	decrypted := make([]byte, len(encrypted))
//	for bs, be := 0, cipher.BlockSize(); bs < len(encrypted); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
//		cipher.Decrypt(decrypted[bs:be], encrypted[bs:be])
//	}
//
//	trim := 0
//	if len(decrypted) > 0 {
//		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
//	}
//
//	//return string(decrypted[:trim]), nil
//	da, err := Base64Decode(string(decrypted[:trim]))
//	return string(da), err
//}
//
//func main() {
//	key := "iscConfigService"
//	plaintext := "sdf3"
//
//	fmt.Printf("原    文：%s\n", plaintext)
//	encrypted := aesEncryptECB(plaintext, key)
//	fmt.Printf("加密结果：%s\n", encrypted)
//	decrypted, err := aesDecryptECB(encrypted, key)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Printf("解密结果：%s\n", decrypted)
//}
//
//
//
//func padding(src []byte, blocksize int) []byte {
//	padnum := blocksize - len(src)%blocksize
//	pad := bytes.Repeat([]byte{byte(padnum)}, padnum)
//	return append(src, pad...)
//}
//
//func unpadding(src []byte) []byte {
//	n := len(src)
//	unpadnum := int(src[n-1])
//	return src[:n-unpadnum]
//}
//
//func EncryptAES2(src []byte, key []byte) []byte {
//	block, _ := aes.NewCipher(key)
//	//src = padding(src, block.BlockSize())
//	blockmode := cipher.NewCBCEncrypter(block, key)
//	blockmode.CryptBlocks(src, src)
//	return src
//}
//
//func DecryptAES2(src []byte, key []byte) []byte {
//	block, _ := aes.NewCipher(key)
//	blockmode := cipher.NewCBCDecrypter(block, key)
//	blockmode.CryptBlocks(src, src)
//	src = unpadding(src)
//	return src
//}
//
//func decodeMsg(result, key string) string {
//	middle, err := Base64Decode(result)
//	if err != nil {
//		fmt.Println("err1: ", err.Error())
//		return ""
//	}
//	re := DecryptAES2(middle, []byte(key))
//	return string(re)
//}
//
//func encodeMsg(content, key string) string {
//	re := EncryptAES2([]byte(content), []byte(key))
//	return Base64Encode(re)
//}
//
//func Base64Encode(src []byte) string {
//	return base64.StdEncoding.EncodeToString(src)
//}
//
//func Base64Decode(dst string) ([]byte, error) {
//	return base64.StdEncoding.DecodeString(dst)
//}

//key := []byte("iscConfigService")
//plaintext := []byte("sdf3")
//block, err := aes.NewCipher(key)
//if err != nil {
//	panic(err.Error())
//}
//nonce := make([]byte, NONCESIZE)
//aesgcm, err := cipher.NewGCM(block)
//if err != nil {
//	panic(err.Error())
//}
//ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)
//fmt.Println("Encrypted Text is : ", base64.StdEncoding.EncodeToString(ciphertext))

const (
	sigScript  string = "3046022100E4C617A5626E408938EEAC33D7BC0ACD645D8FF0B9A29FE55EC3F9BAB01A07C9022100D0F9E8A545E997F3E431C5E3DADA0EC5F53ACE6A7E9598FD6F3B85BE055D3E10"
	unsignedTx string = "bqh"
	pubScr     string = "MFYwEAYHKoZIzj0CAQYFK4EEAAoDQgAEnpwnu94soZ/z9QdfTEwJAY+Qhz34pLsNFZA/OgmhekTIBBdJyxpHs4GaId3wWrq11OyqmJlu2qCLDtZcGFBjdA=="
	priScr     string = "MIGNAgEAMBAGByqGSM49AgEGBSuBBAAKBHYwdAIBAQQg49U3dRtNk+WZ6E1BJKIM8oHP2zm766HFTF9jjP1RM9qgBwYFK4EEAAqhRANCAASwYY8XyIPfiNnfkC25axj/Lw6Eq64hVaz34PKquSwiujCk3tB21VivJwrUJkXhkCmAxYsYhF6izcNLpOAvDrPO"
)

func main() {
	// java -jar /Users/zhouzhenyong/project/isyscore/isc-config-server/isc-encrypt/target/isc-encrypt-1.0.0.jar 0 asdf iscConfigService
	cmd := exec.Command("java", "-jar", "/Users/zhouzhenyong/file/isyscore/toGo/isc-encrypt-1.0.0.jar", "1", "bfWp5JjIf2TcFSteURVgqg==", "iscConfigService")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(string(out))
}
