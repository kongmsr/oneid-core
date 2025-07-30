package utilo

import (
	"errors"
	"github.com/alice52/jasypt-go"
	"os"
)

var Jasypt = jasypt.New()

func MustEnvStr(key string) string {
	str, err := GetEnvStr(key)
	if err != nil || str == nil {
		return ""
	}

	return *str
}

func GetEnvStr(key string, defVal ...string) (*string, error) {
	val := os.Getenv(key)
	if len(val) == 0 {
		if len(defVal) > 0 {
			return &defVal[0], nil
		}

		return nil, errors.New("env not found")
	}

	decryptVal, err := Jasypt.DecryptWrapper(val)
	return &decryptVal, err
}

func MustEnvStrV2(key string) string {
	val, err := GetEnvStrV2(key)
	if err != nil {
		return ""
	}

	return *val
}

func GetEnvStrV2(key string, defVal ...string) (*string, error) {
	return GetEnv[string](key, defVal...)
}

func MustEnv[T Basic](key string, defVal ...T) T {
	val, err := GetEnv[T](key, defVal...)
	if err != nil {
		var zero T // 用于返回默认零值
		return zero
	}

	return *val
}

func GetEnv[T Basic](key string, defVal ...T) (*T, error) {
	val := os.Getenv(key)
	if len(val) == 0 {
		if len(defVal) > 0 {
			return &defVal[0], nil
		}
		return nil, errors.New("env not found")
	}

	// 解密环境变量
	plainText, err := Jasypt.DecryptWrapper(val)
	if err != nil {
		return nil, err
	}

	return Str2BasicPtr[T](plainText)
}

func ExistEnv(key string) bool {
	val := os.Getenv(key)
	return len(val) > 0
}
