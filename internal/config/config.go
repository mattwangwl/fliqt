package config

import (
	"log"
	"path"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

type Config struct {
	MySQL struct {
		Host     string `mapstructure:"host"`
		Port     string `mapstructure:"port"`
		DBName   string `mapstructure:"dbName"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
	} `mapstructure:"mysql"`
	Redis struct {
		Host     string `mapstructure:"host"`
		Port     string `mapstructure:"port"`
		Password string `mapstructure:"password"`
		DB       int    `mapstructure:"db"`
	} `mapstructure:"redis"`
}

func New() *Config {
	impl := &Config{}
	impl.Load()
	return impl
}

func (c *Config) Load() {
	// 設置 Viper 讀取的配置文件名稱和路徑
	viper.SetConfigName("config")                         // 配置文件名稱 (不包括擴展名)
	viper.SetConfigType("yaml")                           // 配置文件類型
	viper.AddConfigPath(path.Join(c.rootDir(), "config")) // 配置文件路徑

	// 讀取配置文件
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	// 將配置文件中的內容加載到 Config 結構中
	if err := viper.Unmarshal(c); err != nil {
		log.Fatalf("Failed to unmarshal config: %v", err)
	}
}

func (c *Config) rootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return path.Join(filepath.Dir(d), "..")
}
