package setting

type Config struct {
	Server ServerSetting `mapstructure:"server"`
	MySQL  MySQLSetting  `mapstructure:"mysql"`
	Logger LoggerSetting `mapstructure:"logger"`
	Redis  RedisSetting  `mapstructure:"redis"`
}

type ServerSetting struct {
	Host       string `mapstructure:"host"`
	Port       int    `mapstructure:"port"`
	Enviroment string `mapstructure:"enviroment"`
}

type MySQLSetting struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
	Dbname       string `mapstructure:"database"`
	MaxIdleConns int    `mapstructure:"max_idle_cons"`
	MaxOpenCons  int    `mapstructure:"max_open_cons"`
	ConnMaxTime  int    `mapstructure:"con_max_life_time"`
}

type RedisSetting struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Database int    `mapstructure:"database"`
	PoolSize int    `mapstructure:"pool_size"`
}

type LoggerSetting struct {
	FileName   string `mapstructure:"file_name"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	Compress   bool   `mapstructure:"compress"`
	LogLevel   string `mapstructure:"log_level"`
	MaxBackups int    `mapstructure:"max_backups"`
}
