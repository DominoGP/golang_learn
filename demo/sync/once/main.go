package main

import (
	"fmt"
	"sync"
)

type Config struct {
	config map[string]string
}

var (
	config *Config
	once   sync.Once
)

func GetConfig() *Config {
	once.Do(func() {
		fmt.Println("init config...")
		config = &Config{
			config: map[string]string{
				"c1": "v1",
				"c2": "v2",
			},
		}
	})
	return config
}

func main() {
	// 第一次需要获取配置信息，初始化 config
	cfg := GetConfig()
	fmt.Println("c1: ", cfg.config["c1"])

	// 第二次需要，此时 config 已经被初始化过，无需再次初始化
	cfg2 := GetConfig()
	fmt.Println("c2: ", cfg2.config["c2"])
}
