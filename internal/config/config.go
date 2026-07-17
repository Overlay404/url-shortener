package config

type Config struct {
	App   AppConfig   `yaml:"app"`
	Redis RedisConfig `yaml:"redis"`
}

type AppConfig struct {
	LevelLogs string `yaml:"level_logs"`
}

type RedisConfig struct {
	Addr     string `yaml:"addr"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}
