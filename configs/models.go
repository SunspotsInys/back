package configs

type AppConfig struct {
	Name           string `mapstructure:"name"`
	Mode           string `mapstructure:"mode"`
	Version        string `mapstructure:"version"`
	Port           int    `mapstructure:"port"`
	*Admin         `mapstructure:"admin"`
	*MariaDBConfig `mapstructure:"mariadb"`
	*LogConfig     `mapstructure:"log"`
}

type Admin struct {
	AdminUsername string `mapstructure:"name"`
	AdminPassword string `mapstructure:"passwd"`
	AdminAvatar   string `mapstructure:"avatar"`
}

type MariaDBConfig struct {
	DBHost         string `mapstructure:"host"`
	DBUser         string `mapstructure:"user"`
	DBPasswd       string `mapstructure:"password"`
	DBName         string `mapstructure:"db_name"`
	DBPort         int    `mapstructure:"port"`
	DBMaxOpenConns int    `mapstructure:"max_open_conns"`
	DBMaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type LogConfig struct {
	LogLevel      string `mapstructure:"level"`
	LogFilename   string `mapstructure:"filename"`
	LogMaxSize    int    `mapstructure:"max_size"`
	LogMaxAge     int    `mapstructure:"max_age"`
	LogMaxBackups int    `mapstructure:"max_backups"`
	LogPrefix     string `mapstructure:"prefix"`
	LogConsole    bool   `mapstructure:"console"`
}
