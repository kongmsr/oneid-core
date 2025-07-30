package utilo

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToPrt(t *testing.T) {
	fmt.Println(ToPrt(1))
	fmt.Println(ToPrt("1"))
	fmt.Println(ToPrt(true))
	fmt.Println(ToPrt([]int{1, 2, 3}))
	fmt.Println(ToPrt(map[string]int{"a": 1, "b": 2}))
	fmt.Println(ToPrt(func() {}))
	fmt.Println(ToPrt(struct{}{}))
}

func TestToVl(t *testing.T) {
	var a *int64
	fmt.Println(ToVl(a))
	var b *string
	fmt.Println(ToVl(b))
	fmt.Println(ToVl(ToPrt(1)))
	fmt.Println(ToVl(ToPrt(true)))
	fmt.Println(ToVl(ToPrt([]int{1, 2, 3})))
	fmt.Println(ToVl(ToPrt(map[string]int{"a": 1, "b": 2})))
	//fmt.Println(ToVl(ToPrt(func() {})))
	fmt.Println(ToVl(ToPrt(struct{}{})))
}

// 每种类型的成功测试
func TestStr2Basic_Bool(t *testing.T) {
	v, err := Str2Basic[bool]("true")
	assert.NoError(t, err)
	assert.Equal(t, true, v)
}

func TestStr2Basic_Int(t *testing.T) {
	v, err := Str2Basic[int]("123")
	assert.NoError(t, err)
	assert.Equal(t, 123, v)
}

func TestStr2Basic_Int8(t *testing.T) {
	v, err := Str2Basic[int8]("127")
	assert.NoError(t, err)
	assert.Equal(t, int8(127), v)
}

func TestStr2Basic_Int16(t *testing.T) {
	v, err := Str2Basic[int16]("32000")
	assert.NoError(t, err)
	assert.Equal(t, int16(32000), v)
}

func TestStr2Basic_Int32(t *testing.T) {
	v, err := Str2Basic[int32]("2147483647")
	assert.NoError(t, err)
	assert.Equal(t, int32(2147483647), v)
}

func TestStr2Basic_Int64(t *testing.T) {
	v, err := Str2Basic[int64]("9223372036854775807")
	assert.NoError(t, err)
	assert.Equal(t, int64(9223372036854775807), v)
}

func TestStr2Basic_Uint(t *testing.T) {
	v, err := Str2Basic[uint]("123")
	assert.NoError(t, err)
	assert.Equal(t, uint(123), v)
}

func TestStr2Basic_Uint8(t *testing.T) {
	v, err := Str2Basic[uint8]("255")
	assert.NoError(t, err)
	assert.Equal(t, uint8(255), v)
}

func TestStr2Basic_Uint16(t *testing.T) {
	v, err := Str2Basic[uint16]("65535")
	assert.NoError(t, err)
	assert.Equal(t, uint16(65535), v)
}

func TestStr2Basic_Uint32(t *testing.T) {
	v, err := Str2Basic[uint32]("4294967295")
	assert.NoError(t, err)
	assert.Equal(t, uint32(4294967295), v)
}

func TestStr2Basic_Uint64(t *testing.T) {
	v, err := Str2Basic[uint64]("18446744073709551615")
	assert.NoError(t, err)
	assert.Equal(t, uint64(18446744073709551615), v)
}

func TestStr2Basic_Float32(t *testing.T) {
	v, err := Str2Basic[float32]("3.14")
	assert.NoError(t, err)
	assert.InDelta(t, float32(3.14), v, 0.0001)
}

func TestStr2Basic_Float64(t *testing.T) {
	v, err := Str2Basic[float64]("3.1415926")
	assert.NoError(t, err)
	assert.InDelta(t, float64(3.1415926), v, 0.0001)
}

func TestStr2Basic_String(t *testing.T) {
	v, err := Str2Basic[string]("hello")
	assert.NoError(t, err)
	assert.Equal(t, "hello", v)
}

// 错误输入测试
func TestStr2Basic_Bool_Error(t *testing.T) {
	_, err := Str2Basic[bool]("notabool")
	assert.Error(t, err)
}

func TestStr2Basic_Int_Error(t *testing.T) {
	_, err := Str2Basic[int]("abc")
	assert.Error(t, err)
}

// ✅ 编译错误示例（不可测试但说明约束生效）
//func TestStr2Basic_Unsupported(t *testing.T) {
//	_, err := Str2Basic[time.Time]("2023-01-01")
//	assert.Error(t, err) // 编译期就报错：T does not satisfy Convertible
//}
