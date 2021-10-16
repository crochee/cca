package config

import (
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

// LoadConfig init Config
func LoadConfig(path string) error {
	pathList := strings.Split(path, "/")
	lastIndex := len(pathList) - 1
	index := strings.LastIndexByte(pathList[lastIndex], '.')
	if index == -1 {
		return errors.New("path hasn't ext")
	}

	viper.AddConfigPath(path[:len(path)-len(pathList[lastIndex])])            // 设置配置文件路径
	viper.SetConfigName(pathList[lastIndex][:index])                          // 设置配置文件名称
	viper.SetConfigType(strings.TrimPrefix(pathList[lastIndex][index:], ".")) // 设置配置文件类型

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("read config failed.%w", err)
	}
	return nil
}
