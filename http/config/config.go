package config

import "fmt"

// 设置为私有变量，防止被篡改
var config = new(Config)

func C() *Config {
	return config
}

// Config 应用配置
type Config struct {
	Apps  *App   `mapstructure:"server"`
	Log   *Log   `mapstructure:"log"`
	MySQL *MySQL `mapstructure:"mysql" `
	Oss   *Oss   `mapstructure:"oss"`
}

type App struct {
	HTTP        *Server `mapstructure:"http"`
	Interaction *Server `mapstructure:"interaction"`
	Relation    *Server `mapstructure:"relation"`
}

type Server struct {
	Name string `mapstructure:"name"`
	Mode string `mapstructure:"mode"`
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

func (s *Server) GetAddr() string {
	return fmt.Sprintf("%s:%s", s.Host, s.Port)
}

// Log todo
type Log struct {
	Level        string `mapstructure:"level"`
	InfoFilename string `mapstructure:"infoFilename"`
	ErrFilename  string `mapstructure:"errFilename"`
	MaxSize      int    `mapstructure:"max_size"`
	MaxAge       int    `mapstructure:"max_age"`
	MaxBackups   int    `mapstructure:"max_backups"`
}

// MySQL todo
type MySQL struct {
	Host         string `mapstructure:"host"`
	Port         string `mapstructure:"port"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	Dbname       string `mapstructure:"dbname"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
}

type Oss struct {
	AccessKeyId     string `mapstructure:"ALI_AK"`
	AccessKeySecret string `mapstructure:"ALI_SK"`
	OssEndpoint     string `mapstructure:"ALI_OSS_ENDPOINT"`
	BucketName      string `mapstructure:"ALI_BUCKET_NAME"`
	OssVideoDir     string `mapstructure:"	OSS_VIDEO_DIR"`
	PlayUrlPrefix   string `mapstructure:"PLAY_URL_PREFIX"`
}

//type RedisConfig struct {
//	Host     string `mapstructure:"host"`
//	Port     int    `mapstructure:"port"`
//	Password string `mapstructure:"password"`
//	Db       int    `mapstructure:"db"`
//}
