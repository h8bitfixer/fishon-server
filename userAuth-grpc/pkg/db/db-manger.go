package db

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	mysqlDB *gorm.DB
	redisDB *redis.Client
)

// InitializeMySQL initializes the MySQL database connection.
func InitializeMySQL() error {
	dsn := "root:sandman@tcp(mysql:3306)/?" // Replace with your MySQL connection details
	var err error

	mysqlDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to MySQL database: %v", err)
	}

	// Perform additional MySQL-specific setup or configurations

	return nil
}

// GetMySQLDB returns the MySQL GORM instance.
func GetMySQLDB() *gorm.DB {
	return mysqlDB
}

// InitializeRedis initializes the Redis database connection.
func InitializeRedis() error {
	redisDB = redis.NewClient(&redis.Options{
		Addr:     "redis:6379", // Replace with your Redis server address
		Password: "123456",     // Replace with your Redis password, if any
		DB:       0,            // Replace with your Redis database number
	})

	// Ping the Redis server to test the connection
	err := redisDB.Ping(context.TODO()).Err()
	if err != nil {
		fmt.Printf("Failed to connect to Redis: %v\n", err)
		return err
	}

	return nil
}

// GetRedisDB returns the Redis GORM instance.
func GetRedisDB() *redis.Client {
	return redisDB
}
