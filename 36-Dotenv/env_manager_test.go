package main

import (
	"os"
	"testing"
)

func TestEnvManager_GetEnv(t *testing.T) {
	// 设置测试用例。
	testCases := []struct {
		name     string // 测试用例名称。
		key      string // 环境变量名称。
		fallback string // 默认值。
		value    string // 环境变量值。
		expected string // 预期结果。
	}{
		{
			name:     "key exists",
			key:      "S3_BUCKET",
			fallback: "default_bucket",
			value:    "my_bucket",
			expected: "my_bucket",
		},
		{
			name:     "key does not exist",
			key:      "SECRET_KEY",
			fallback: "default_secret_key",
			value:    "",
			expected: "default_secret_key",
		},
	}

	for _, tc := range testCases {
		// 如果tc.value 為空，則不設置環境變數。
		if tc.value != "" {
			os.Setenv(tc.key, tc.value)
		}

		env := EnvManager{}

		actual := env.GetEnv(tc.key, tc.fallback)

		// 检查结果是否符合预期。
		if actual != tc.expected {
			t.Errorf("%s: expected %s, but got %s", tc.name, tc.expected, actual)
		}

		// 清除环境变量。
		os.Unsetenv(tc.key)
	}
}

func TestEnvManager_LoadEnv(t *testing.T) {
	// 创建 EnvManager 实例。
	env := EnvManager{}

	// 调用 LoadEnv 方法。
	err := env.LoadEnv()

	// 检查结果是否符合预期。
	if err != nil {
		t.Errorf("expected no error, but got %v", err)
	}
}
