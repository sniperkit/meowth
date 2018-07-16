package lib

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"time"
)

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

func GenerateRandomString(s int) string {
	b, _ := GenerateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b)
}

func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func EncodeSha256(input string) string {
	sha_256 := sha256.New()
	sha_256.Write([]byte(input))
	sha := sha_256.Sum(nil)
	return fmt.Sprintf("%s", sha)
}

func EncodeSha512(input string) string {
	sha_512 := sha512.New()
	sha_512.Write([]byte(input))
	sha := sha_512.Sum(nil)
	return fmt.Sprintf("%s", sha)
}

func EncodeSha512_256(input string) string {
	sha_512 := sha512.New()
	sha512.Sum512_256([]byte(input))
	sha := sha_512.Sum(nil)
	return fmt.Sprintf("%s", sha)
}

func EncodeHmac512(input string) string {
	hmac := hmac.New(sha512.New, []byte(input))
	hmac.Write([]byte(input))
	return base64.StdEncoding.EncodeToString(hmac.Sum(nil))
}

func RandomSha512(length int) string {
	code := fmt.Sprintf("%v+%d", GenerateRandomString(length), makeTimestamp())
	return EncodeSha512(code)
}

func RandomSha256(length int) string {
	code := fmt.Sprintf("%v+%d", GenerateRandomString(length), makeTimestamp())
	return EncodeSha256(code)
}

func RandomSha512_256(length int) string {
	code := fmt.Sprintf("%v+%d", GenerateRandomString(length), makeTimestamp())
	return EncodeSha512_256(code)
}

func RandomHmac512(length int) string {
	code := fmt.Sprintf("%v+%d", GenerateRandomString(length), makeTimestamp())
	return EncodeHmac512(code)
}
