package transport

import (
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
)

func Signature(apiKey, requestURI string) (string, error) {
	mac := hmac.New(sha1.New, []byte(apiKey))
	_, err := mac.Write([]byte(requestURI))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", mac.Sum(nil)), nil
}
