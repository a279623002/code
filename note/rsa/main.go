package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"os"
)

var (
	privateKeyPath = "./pem/private_ssl.pem"
	publicKeyPath  = "./pem/public_ssl.pem"
)

// 读取PKCS1格式私钥
func ReadRSAPKCS1PrivateKey(path string) (*rsa.PrivateKey, error) {
	// 读取文件
	context, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	// pem解码
	pemBlock, _ := pem.Decode(context)
	// x509解码
	privateKey, err := x509.ParsePKCS1PrivateKey(pemBlock.Bytes)
	return privateKey, err
}

// 读取公钥(包含PKCS1和PKCS8)
func ReadRSAPublicKey(path string) (*rsa.PublicKey, error) {
	var err error
	// 读取文件
	readFile, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	// 使用pem解码
	pemBlock, _ := pem.Decode(readFile)
	var pkixPublicKey interface{}
	if pemBlock.Type == "RSA PUBLIC KEY" {
		// -----BEGIN RSA PUBLIC KEY-----
		pkixPublicKey, err = x509.ParsePKCS1PublicKey(pemBlock.Bytes)
	} else if pemBlock.Type == "PUBLIC KEY" {
		// -----BEGIN PUBLIC KEY-----
		pkixPublicKey, err = x509.ParsePKIXPublicKey(pemBlock.Bytes)
	}
	if err != nil {
		return nil, err
	}
	publicKey := pkixPublicKey.(*rsa.PublicKey)
	return publicKey, nil
}

func ReadKey() {
	privateKey, err := ReadRSAPKCS1PrivateKey(privateKeyPath)
	if err != nil {
		panic(err)
	}
	fmt.Println(privateKey)
	publicKey, err := ReadRSAPublicKey(publicKeyPath)
	if err != nil {
		panic(err)
	}
	fmt.Println(publicKey)
}

// 加密(使用公钥加密)
func RSAEncrypt(data, publicKeyPath string) (string, error) {
	// 获取公钥
	// ReadRSAPublicKey代码在 【3.读取密钥】
	rsaPublicKey, err := ReadRSAPublicKey(publicKeyPath)
	if err != nil {
		return "", err
	}
	// 加密
	encryptPKCS1v15, err := rsa.EncryptPKCS1v15(rand.Reader, rsaPublicKey, []byte(data))
	if err != nil {
		return "", err
	}
	// 把加密结果转成Base64
	encryptString := base64.StdEncoding.EncodeToString(encryptPKCS1v15)
	return encryptString, err
}

// 解密(使用私钥解密)
func RSADecrypt(base64data, privateKeyPath string) (string, error) {
	// data反解base64
	decodeString, err := base64.StdEncoding.DecodeString(base64data)
	if err != nil {
		return "", err
	}
	// 读取密钥
	rsaPrivateKey, err := ReadRSAPKCS1PrivateKey(privateKeyPath)
	if err != nil {
		return "", err
	}
	// 解密
	decryptPKCS1v15, err := rsa.DecryptPKCS1v15(rand.Reader, rsaPrivateKey, decodeString)
	return string(decryptPKCS1v15), err
}

// 对数据进行数字签名
func GetRSASign(data, privateKeyPath string) (string, error) {
	// 读取私钥
	privateKey, err := ReadRSAPKCS1PrivateKey(privateKeyPath)
	if err != nil {
		return "", err
	}
	// 计算Sha1散列值
	hash := sha256.New()
	hash.Write([]byte(data))
	sum := hash.Sum(nil)
	// 从1.5版本规定，使用RSASSA-PKCS1-V1_5-SIGN 方案计算签名
	signPKCS1v15, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, sum)
	// 结果转成base64
	toString := base64.StdEncoding.EncodeToString(signPKCS1v15)
	return toString, err
}

// 验证签名
func VerifyRsaSign(data, publicKeyPath, base64Sign string) (bool, error) {
	// 反解base64
	sign, err := base64.StdEncoding.DecodeString(base64Sign)
	if err != nil {
		return false, err
	}
	// 获取公钥
	publicKey, err := ReadRSAPublicKey(publicKeyPath)
	if err != nil {
		return false, err
	}
	// 计算Sha1散列值
	hash := sha256.New()
	hash.Write([]byte(data))
	bytes := hash.Sum(nil)
	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, bytes, sign)
	return err == nil, err
}

func main() {
	eStr, err := RSAEncrypt("hi, zzq", publicKeyPath)
	if err != nil {
		panic(err)
	}
	fmt.Println(eStr)
	dStr, err := RSADecrypt(eStr, privateKeyPath)
	if err != nil {
		panic(err)
	}
	fmt.Println(dStr)

}
