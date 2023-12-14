package utils

import (
	"encoding/base64"
	"encoding/hex"
)

func UTF8ToHex(str string) string {
	encoded := hex.EncodeToString([]byte(str))
	return encoded
}

func HexToUTF8(hexStr string) (string, error) {
	decoded, err := hex.DecodeString(hexStr)
	if nil != err {
		return "", err
	}

	return string(decoded), nil
}

func StringToBase64(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func Base64ToString(str string) (string, error) {
	decodeString, err := base64.StdEncoding.DecodeString(str)
	if nil != err {
		return "", err
	}

	return string(decodeString), nil
}

func StringToHex(input string) string {
	return hex.EncodeToString([]byte(input))
}

func HexToBytes(str string) []byte {
	h, _ := hex.DecodeString(str)
	return h
}
