package config

type Config struct {
	App   AppConfig   `yaml:"app"`
	Redis RedisConfig `yaml:"redis"`
}

type AppConfig struct {
	LevelLogs string `yaml:"level_logs"`
}

type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DB       string `yaml:"db"`
}
