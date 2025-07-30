package utilo

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func init() {
	_ = os.Setenv("ENV_STR", "test")
	_ = os.Setenv("ENV_BOOL", "true")
	_ = os.Setenv("ENV_INT", "11")
	_ = os.Setenv("ENV_UNIT", "1")
	_ = os.Setenv("ENV_FLOAT", "3.14")
}

func TestGetEnvStr(t *testing.T) {
	fmt.Println(MustEnvStr("ENV_STR"))
	assert.Equal(t, MustEnvStr("ENV_STR"), "test")
	fmt.Println(MustEnvStr("ENV_STR_NO"))
	assert.Equal(t, MustEnvStr("ENV_STR_NO"), "")

	fmt.Println(MustEnvStrV2("ENV_STR"))
	assert.Equal(t, MustEnvStrV2("ENV_STR"), "test")
	fmt.Println(MustEnvStrV2("ENV_STR_NO"))
	assert.Equal(t, MustEnvStrV2("ENV_STR_NO"), "")
}

func TestGetEnv(t *testing.T) {
	fmt.Println(MustEnv[bool]("ENV_STR"))
	fmt.Println(MustEnv[bool]("ENV_STR_NO"))

	fmt.Println(MustEnv[bool]("ENV_BOOL"))
	fmt.Println(MustEnv[bool]("ENV_BOOL_NO"))
}
