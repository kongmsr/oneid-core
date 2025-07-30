package utilo

import (
	"errors"
	"github.com/spf13/cast"
	"reflect"
)

func ToPrt[T any](i T) *T {
	return &i
}

// ToVl 返回指针i指向的值，如果i是nil，则返回T的零值。
func ToVl[T any](i *T) T {
	if i == nil {
		var zero T
		return zero
	}
	return *i
}

type Basic interface {
	~bool |
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
	~float32 | ~float64 |
	~string
}

func Str2BasicPtr[T Basic](str string) (*T, error) {
	basic, err := Str2Basic[T](str)
	if err != nil {
		return nil, err
	}
	return &basic, nil
}

func Str2Basic[T Basic](str string) (T, error) {
	var result any
	var zero T // 用于返回默认零值

	switch any(zero).(type) {
	case bool:
		if res, err := cast.ToBoolE(str); err != nil {
			return zero, err
		} else {
			result = res
		}
	case int:
		if res, err := cast.ToIntE(str); err != nil {
			return zero, err
		} else {
			result = res
		}
	case int8:
		if res, err := cast.ToInt8E(str); err != nil {
			return zero, err
		} else {
			result = res
		}
	case int16:
		if res, err := cast.ToInt16E(str); err != nil {
			return zero, err
		} else {
			result = res
		}
	case int32:
		if res, err := cast.ToInt32E(str); err != nil {
			return zero, err
		} else {
			result = res
		}
	case int64:
		if res, err := cast.ToInt64E(str); err != nil {
			return zero, err
		} else {
			result = res
		}

	case uint:
		if res, err := cast.ToUintE(str); err != nil {
			return zero, err
		} else {
			result = res
		}
	case uint8:
		if res, err := cast.ToUint8E(str); err != nil {
			return zero, err
		} else {
			result = res
		}
	case uint16:
		if res, err := cast.ToUint16E(str); err != nil {
			return zero, err
		} else {
			result = res
		}
	case uint32:
		if res, err := cast.ToUint32E(str); err != nil {
			return zero, err
		} else {
			result = res
		}
	case uint64:
		if res, err := cast.ToUint64E(str); err != nil {
			return zero, err
		} else {
			result = res
		}

	case float32:
		result = cast.ToFloat32(str)
		if res, err := cast.ToFloat32E(str); err != nil {
			return zero, err
		} else {
			result = res
		}
	case float64:
		if res, err := cast.ToFloat64E(str); err != nil {
			return zero, err
		} else {
			result = res
		}

	case string:
		result = str
	default:
		return zero, errors.New("unsupported type: " + reflect.TypeOf(*new(T)).String())
	}

	// 将转换后的值赋给变量
	casted := result.(T)
	return casted, nil
}
