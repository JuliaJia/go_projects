package conf

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/infraboard/mcube/logger/zap"
	"sync"
	"time"
)

//配置通过对象来进行隐射
//我们需要定义配置对象的数据结构

var global *Config

var db *sql.DB

func C() *Config {
	if global != nil {
		panic("config required!")
	}
	return global
}

func SetGlobalConfig(conf *Config) {
	global = conf
}

func NewDefaultConfig() *Config {
	return &Config{
		App:   newDefaultApp(),
		Mysql: newDefaultMysql(),
		Log:   newDefaultLog(),
	}
}

type Config struct {
	App   *app
	Mysql *mysql
	Log   *log
}

func newDefaultApp() *app {
	return &app{
		Name: "demo",
		Host: "localhost",
		Port: "8050",
		Key:  "default",
	}
}

type app struct {
	Name string `toml:"name"`
	Port string `toml:"port"`
	Host string `toml:"host"`
	Key  string `toml:"key"`
}

func newDefaultMysql() *mysql {
	return &mysql{
		Host:        "localhost",
		Port:        "3306",
		Username:    "root",
		Password:    "123456",
		Database:    "go_learn",
		MaxOpenConn: 100,
		MaxLifeTime: 10 * 60 * 60,
		MaxIdleConn: 20,
		MaxIdleTime: 5 * 60 * 60,
	}
}

type mysql struct {
	Host        string `toml:"host" env:"MYSQL_HOST"`
	Port        string `toml:"port" env:"MYSQL_PORT"`
	Username    string `toml:"username" env:"MYSQL_USERNAME"`
	Password    string `toml:"password" env:"MYSQL_PASSWORD"`
	Database    string `toml:"database" env:"MYSQL_DATABASE"`
	MaxOpenConn int    `toml:"max_open_conn" env:"D_MYSQL_MAX_OPEN_CONN"`
	MaxLifeTime int    `toml:"max_life_time" env:"D_MYSQL_MAX_LIFE_TIME"`
	MaxIdleConn int    `toml:"max_open_conn" env:"D_MYSQL_MAX_IDLE_CONN"`
	MaxIdleTime int    `toml:"max_open_conn" env:"D_MYSQL_MAX_IDLE_TIME"`
	lock        sync.Mutex
}

func (m *mysql) GetDBConn() (*sql.DB, error) {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&multiStatements=true", m.Username, m.Password, m.Host, m.Port, m.Database)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("connect to mysql<%s> error,%s", dsn, err.Error())
	}
	db.SetMaxOpenConns(m.MaxOpenConn)
	db.SetMaxIdleConns(m.MaxIdleConn)
	db.SetConnMaxLifetime(time.Second * time.Duration(m.MaxLifeTime))
	db.SetConnMaxIdleTime(time.Second * time.Duration(m.MaxIdleTime))
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("ping mysql<%s> error, %s", dsn, err.Error())
	}
	return db, nil
}

func (m *mysql) GetDB() (*sql.DB, error) {
	m.lock.Lock()
	defer m.lock.Unlock()
	if db == nil {
		conn, err := m.GetDBConn()
		if err != nil {
			return nil, err
		}
		db = conn
	}
	return db, nil
}

func newDefaultLog() *log {
	return &log{
		Level:  zap.DebugLevel.String(),
		Format: TextFormat,
		Path:   "logs",
		To:     ToStdout,
	}
}

type log struct {
	Level  string    `toml:"level" env:"LOG_LEVEL"`
	Format LogFormat `toml:"format" env:"LOG_FORMAT"`
	Path   string    `toml:"path" env:"LOG_PATH_DIR"`
	To     LogTo     `toml:"to" env:"LOG_TO"`
}
