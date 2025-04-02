package config

//Config holds application configuration
type Config struct {
	DBUser string
	DBPassword string
	DBName string
	DBHost string
	DBPort string

}

func LoadConfig()*Config{
	config := &Config{
        DBUser:     getEnv("DB_USER", "postgres"),
        DBPassword: getEnv("DB_PASSWORD", "postgres"),
        DBName:     getEnv("DB_NAME", "gotasker"),
        DBHost:     getEnv("DB_HOST", "localhost"),
        Port:       getEnv("PORT", "8080"),
    }
	return config

}