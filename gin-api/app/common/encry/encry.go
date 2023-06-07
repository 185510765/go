package encry

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"fmt"

	"github.com/smartwalle/crypto4go"
)

// rsa 私钥解密 前端公钥加密
func RsaDecrypt(data string) (string, error) {
	privateKey := getPrivateKey()

	priKey, err := crypto4go.ParsePKCS1PrivateKey(crypto4go.FormatPKCS1PrivateKey(string(privateKey)))
	if err != nil {
		priKey, err = crypto4go.ParsePKCS8PrivateKey(crypto4go.FormatPKCS8PrivateKey(string(privateKey)))
		if err != nil {
			fmt.Println("ParsePKCS8PrivateKey : ", err.Error())
			return "", err
		}
	}
	// 转成base64
	key, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		fmt.Println("base64.RawURLEncoding.DecodeString : ", err.Error())
		return "", err
	}
	partLen := priKey.N.BitLen() / 8
	chunks := split([]byte(key), partLen)
	buffer := bytes.NewBufferString("")

	for _, chunk := range chunks {
		decrypted, err := rsa.DecryptPKCS1v15(rand.Reader, priKey, chunk)
		if err != nil {
			return "", err
		}
		buffer.Write(decrypted)
	}
	return buffer.String(), err
}

func getPrivateKey() string {
	return `-----BEGIN PRIVATE KEY-----
	MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCfO0RTJVvE1E/L
	b3QwdP0lq4ItrkgQN7S8RMiWjLyJBJO3qaDM0pTuK/s6V1ovTnGNQe9P/ta4tzez
	l7Ainebk25A5UkelP8CwrrxxEqbAM9It+XqzuGe+ntr5yLW1SieDNlC4+0MJH1Eq
	odEv8GzEHAKDTkYod44q4w0ScqyXIkqbyhIiHeJn0EKMKP44LWhp9vkMj0p3DEUI
	vmUBdyX+RBpqAr76nJnMG91S+G8H5/1XmU3CSgHfoL9ofje2GUMifs1yaQPvPn9C
	ft/6MtVrJ4URcXkIefJmaNqWnRk8xVqrXhR9LW7+MwaLBdKvPTH274xIHk25SMeW
	jAVTtgeBAgMBAAECggEAEIRHTBCd+kbDmCiRds1LzPKDaFWhp/z/RGAmJmemzteo
	su5nfZeCV0o8nwqckjei07I5LqUShMqMfpaHcK/MWYTk6u07/UolpK4bYec7YMh/
	TsGIFshUCpPTxwEoyAtPGe8yZF270GBtPOzYFA7uWVFGeXRPFreACbbPjlBS0sJT
	84jWtu+pi4q+q65+u3aS9Qh/Ci8+NqRK8qLW0QDuKa5plKOgA1FSk1vnm6pb1Tcd
	ViNBrjSk8OOh77j57hze19c9CRTtyJuk+dQFWMqi1RKNi2aMlZae4bJLl98plvLs
	W1VXrI7Q03GA3wTuK15lu5Z00AWnnRPk1mQMUYZjmQKBgQDMDuknb/XBSeM6Xdb9
	jYUinnyYVongqTJzSPJ2LUusmw8/zUbQTjPtYOkoALHEv0ruVrl66yBKtk2uZQvJ
	7uUkLtmLTHvPV1mOUVFbdj0T1FN3xrIPeB/wsze/GzsYodVAYcBeSB/E5finHBrO
	99W+OS+EkV6SYiW4T1VVyNOb+wKBgQDHw01wKjJasEEm+B/ARdW/cgXcWs1wNdfz
	ro6lZr6sbc3spgcmEhxh1Hjd25QvO5+ZPKcxTq15y9XcFyU+IyJ1/ZWPQoxQ7gCZ
	EWApRuEO8zBxFnmE+UwOMMZpM7qE7oUHcuaUeCRIUwmhndj4G0CMJO0hJxptIm8h
	3t4DAk81swKBgEZSSSUvHzkKNoVxu1pLv/rLNVLmV8OVa25xUCVLvM1x7lJlcqbs
	nFaM1CzV+G0+Ixt5xZfHmaxFoQWdiu9/JXZPsuafZ/dvOcyi12+2kpvXyx/22Hwe
	QJuZl3eDcd0uQChcx4d2QYSAYC1usQpsPDu+x1JfKoE105vtsxHKEKqrAoGAZhrj
	iANpfYU9qAeHYyXO3W4QpNMc0tASs5FzhAOCTmxJpz4txT3YmACcTvofQg09xHuG
	EePfM7QGedqyxJH82UZmnbUsN2mPkK2a8z4xZwzSo9sS/e7W+yHfKIKyQaQRBv+p
	8BWAph8beypNbnOviHKPajSsgz/oge2CmRMjHZUCgYEAttZWnJtUW7cOu7AvEP+T
	fLkqeQ8slZ6P9pgHBZBX+S6Li7zwh74RPMYYLUGBH6U4bU/pz8IBVUG0OD5IPxhb
	Qvruaj3nNzvT++g82/9jqLASp/sMxV0xt/9YstT1GS6j8NQcXQ2eldsnRHpZ2PeC
	frpHtnUD8xO4XoNmdS/hFfI=
	-----END PRIVATE KEY-----
	`
}

func split(buf []byte, lim int) [][]byte {
	var chunk []byte
	chunks := make([][]byte, 0, len(buf)/lim+1)
	for len(buf) >= lim {
		chunk, buf = buf[:lim], buf[lim:]
		chunks = append(chunks, chunk)
	}
	if len(buf) > 0 {
		chunks = append(chunks, buf[:len(buf)])
	}
	return chunks
}
