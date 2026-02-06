package config

import (
	"os"
)

type Config struct {
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
	ServerPort string `mapstructure:"SERVER_PORT"`
}

// Docker не находит файл .env, viper пока отключен,

// func LoadConfig(path string) (config Config, err error) {
// 	// viper.AddConfigPath(path)
// 	// viper.SetConfigName(".env")
// 	// viper.SetConfigType("env")
// 	// viper.AutomaticEnv()
// 	viper.SetConfigFile(".env")
// 	viper.AutomaticEnv()

// 	err = viper.ReadInConfig()
// 	if err != nil {
// 		log.Println("Error ReadInConfig:", err)
// 		return
// 	}
// 	err = viper.Unmarshal(&config)
// 	if err != nil {
// 		log.Println("Error unmarshalling config:", err)
// 		return
// 	}
// 	return
// }

func LoadConfig() Config {
	return Config{
		DBHost:     getEnv("DB_HOST", "localhost"), // localhost -> db in docker
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "root"),
		DBName:     getEnv("DB_NAME", "todoDB"),
		ServerPort: getEnv("SERVER_PORT", "8080"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
