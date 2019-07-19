package swan_miniprogram

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"errors"
)

func Decrypt(encryptedData, sessionKey, iv, appKey string) (userinfo Userinfo, err error) {
	key, err := base64.StdEncoding.DecodeString(sessionKey)
	if err != nil {
		return
	}

	ivBytes, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return
	}

	encryptedDataBytess, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return
	}

	var block cipher.Block

	if block, err = aes.NewCipher([]byte(key)); err != nil {
		return
	}

	if len(encryptedData) < aes.BlockSize {
		err = errors.New("encryptedData too short")
	}

	cbc := cipher.NewCBCDecrypter(block, ivBytes)
	cbc.CryptBlocks(encryptedDataBytess, encryptedDataBytess)

	encryptedDataBytess = pKCS7UnPadding(encryptedDataBytess)

	if len(encryptedDataBytess) < 20 {
		err = errors.New("bad content")
		return
	}

	// 前面16位可以直接抛弃，17-20表示明文长度
	// TODO
	size := encryptedDataBytess[19]

	if len(encryptedDataBytess) < 20+int(size) {
		err = errors.New("bad content")
		return
	}

	// 最后N位一定是appKey
	if string(encryptedDataBytess[size+20:]) != appKey {
		err = errors.New("illegal appKey")
		return
	}

	err = json.Unmarshal(encryptedDataBytess[20:size+20], &userinfo)
	return
}

func pKCS7UnPadding(plantText []byte) []byte {
	length := len(plantText)
	unPadding := int(plantText[length-1])
	return plantText[:(length - unPadding)]
}
