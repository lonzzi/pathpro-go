package conf

import (
	"fmt"
	"strings"

	"github.com/BurntSushi/toml"
)

type Server struct {
	Host string `toml:"host"` // 主机地址
	Port int    `toml:"port"` // 端口号
}

type Database struct {
	Host     string `toml:"host"`     // 主机地址
	Port     int    `toml:"port"`     // 端口号
	User     string `toml:"user"`     // 用户名
	Password string `toml:"password"` // 密码
	Database string `toml:"database"` // 数据库名
}

type Log struct {
	Level string `toml:"level"` // 日志级别
	File  string `toml:"file"`  // 日志文件
}

type Config struct {
	// 服务配置
	Server Server `toml:"server"`

	// 数据库配置
	Database Database `toml:"database"`

	// 日志配置
	Log Log `toml:"log"`
}

var (
	// 配置文件路径
	configPath string
	// 配置文件
	config *Config
)

// 初始化配置
func Init(path string) error {
	configPath = path
	return loadConfig()
}

// 加载配置
func loadConfig() error {
	// 解析配置文件
	if _, err := toml.DecodeFile(configPath, &config); err != nil {
		return err
	}
	return nil
}

// 获取配置
func GetConfig() *Config {
	return config
}

// 获取服务配置
func GetServerConfig() *Server {
	return &config.Server
}

// 获取数据库配置
func GetDatabaseConfig() *Database {
	return &config.Database
}

// 获取日志配置
func GetLogConfig() *Log {
	return &config.Log
}

func GetString(str string) string {
	strs := strings.Split(str, ".")
	if len(strs) != 2 {
		return ""
	}
	switch strs[0] {
	case "server":
		switch strs[1] {
		case "host":
			return config.Server.Host
		case "port":
			return fmt.Sprintf("%d", config.Server.Port)
		default:
			return ""
		}
	case "database":
		switch strs[1] {
		case "host":
			return config.Database.Host
		case "port":
			return fmt.Sprintf("%d", config.Database.Port)
		case "user":
			return config.Database.User
		case "password":
			return config.Database.Password
		case "database":
			return config.Database.Database
		case "dsn":
			return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
				config.Database.User,
				config.Database.Password,
				config.Database.Host,
				config.Database.Port,
				config.Database.Database,
			)
		default:
			return ""
		}
	case "log":
		switch strs[1] {
		case "level":
			return config.Log.Level
		case "file":
			return config.Log.File
		default:
			return ""
		}
	}
	return ""
}
