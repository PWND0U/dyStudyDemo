package config

type Config struct {
	*App      `json:"app,omitempty" yaml:"app,omitempty"`
	*Database `yaml:"database,omitempty" json:"database,omitempty"`
	*Log      `json:"log,omitempty" yaml:"log,omitempty"`
}
type Database struct {
	Connection string `yaml:"connection,omitempty" json:"connection,omitempty"`
	*Mysql     `yaml:"mysql,omitempty" json:"mysql,omitempty"`
	*Sqlite    `yaml:"sqlite,omitempty" json:"sqlite,omitempty"`
	*Redis     `yaml:"redis,omitempty" json:"redis,omitempty"`
}
type App struct {
	Name     string `json:"name,omitempty" yaml:"name,omitempty"`
	Env      string `json:"env,omitempty" yaml:"env,omitempty"`
	Debug    bool   `json:"debug,omitempty" yaml:"debug,omitempty"`
	AppPort  string `json:"app_port,omitempty" yaml:"appPort,omitempty"`
	AppKey   string `json:"app_key,omitempty" yaml:"appKey,omitempty"`
	AppUrl   string `json:"app_url,omitempty" yaml:"appUrl,omitempty"`
	TimeZone string `json:"time_zone,omitempty" yaml:"timeZone,omitempty"`
}

type Mysql struct {
	Host               string `json:"host,omitempty" yaml:"host,omitempty"`
	DbName             string `json:"dbName,omitempty" yaml:"dbName,omitempty"`
	Port               int    `json:"port,omitempty" yaml:"port,omitempty"`
	Username           string `json:"username,omitempty" yaml:"username,omitempty"`
	Password           string `yaml:"password,omitempty" json:"password,omitempty"`
	Charset            string `json:"charset,omitempty" yaml:"charset,omitempty"`
	MaxIdleConnections int    `json:"maxIdleConnections,omitempty" yaml:"maxIdleConnections,omitempty"`
	MaxOpenConnections int    `json:"maxOpenConnections,omitempty" yaml:"maxOpenConnections,omitempty"`
	MaxLifeSeconds     int    `json:"maxLifeSeconds,omitempty" yaml:"maxLifeSeconds,omitempty"`
}
type Redis struct {
	Host     string `json:"host,omitempty" yaml:"host,omitempty"`
	Port     string `json:"port,omitempty" yaml:"port,omitempty"`
	Password string `json:"password,omitempty" yaml:"password,omitempty"`
	DbName   int    `json:"dbName,omitempty" yaml:"dbName,omitempty"`
}
type Sqlite struct {
	DbName string `json:"dbName,omitempty" yaml:"dbName,omitempty"`
}

type Log struct {
	LogLevel    string `json:"logLevel,omitempty" yaml:"logLevel,omitempty"`
	LogType     string `json:"logType,omitempty" yaml:"logType,omitempty"`
	LogFileName string `json:"logFileName,omitempty" yaml:"logFileName,omitempty"`
	MaxSize     int    `json:"maxSize,omitempty" yaml:"maxSize,omitempty"`
	MaxBackup   int    `json:"maxBackup,omitempty" yaml:"maxBackup,omitempty"`
	MaxAge      int    `json:"maxAge,omitempty" yaml:"maxAge,omitempty"`
	Compress    bool   `json:"compress,omitempty" yaml:"compress,omitempty"`
}

func defaultMysql() *Mysql {
	return &Mysql{
		Host:               "127.0.0.1",
		DbName:             "BlogApi",
		Port:               3306,
		Username:           "blogapi",
		Password:           "12345678",
		Charset:            "utf8mb4",
		MaxIdleConnections: 100,
		MaxOpenConnections: 25,
		MaxLifeSeconds:     300,
	}
}
func defaultRedis() *Redis {
	return &Redis{
		Host:     "127.0.0.1",
		Port:     "6379",
		Password: "",
		DbName:   1,
	}
}
func defaultSqlite() *Sqlite {
	return &Sqlite{
		DbName: "database/database.db",
	}
}
func defaultDatabase() *Database {
	return &Database{
		Connection: "mysql",
		Mysql:      defaultMysql(),
		Redis:      defaultRedis(),
		Sqlite:     defaultSqlite(),
	}
}
func defaultApp() *App {
	a := &App{
		Name:     "BlogApi",
		Env:      "test",
		Debug:    false,
		AppPort:  "10010",
		AppKey:   "568536b0a8e08c239c20cde654dd674d",
		AppUrl:   "",
		TimeZone: "Asia/Shanghai",
	}
	a.AppUrl = "http://localhost:" + a.AppPort
	return a
}

func defaultLog() *Log {
	return &Log{
		LogLevel:    "debug",
		LogType:     "single",
		LogFileName: "storage/logs/logs.log",
		MaxSize:     64,
		MaxBackup:   5,
		MaxAge:      30,
		Compress:    false,
	}
}
