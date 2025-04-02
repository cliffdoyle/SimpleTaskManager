package config


import (
	"fmt"
	"os"
)

//Config holds application configuration
type Config struct {
	DBUser string
	DBPassword string
	DBName string
	DBHost string
	DBPort string

}
//Loads configuration fron environment variables
func LoadConfig()*Config{
	config := &Config{
        DBUser:     getEnv("DB_USER", "postgres"),
        DBPassword: getEnv("DB_PASSWORD", "postgres"),
        DBName:     getEnv("DB_NAME", "gotasker"),
        DBHost:     getEnv("DB_HOST", "localhost"),
        DBPort:       getEnv("PORT", "8080"),
    }
	return config

}
//returns database connection string
func (c *Config)GetDSN()string{
return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	c.DBHost,
	c.DBPort,
	c.DBUser,
	c.DBPassword,
	c.DBName,
)
}

//helper function to get environment variables with a default value
func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}