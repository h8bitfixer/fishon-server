package db

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	config "userAuth-grpc/pkg/common"
)

var (
	mysqlDB *gorm.DB
	redisDB *redis.Client
)

// InitializeMySQL initializes the MySQL database connection.
func InitializeMySQL() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		config.Config.Mysql.DBUserName, config.Config.Mysql.DBPassword, config.Config.Mysql.DBAddress[0], "mysql")
	var err error

	mysqlDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to MySQL database: %v", err)
	}

	// Check the database and table during initialization
	sql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s default charset utf8mb4 COLLATE utf8mb4_unicode_ci;", config.Config.Mysql.DBDatabaseName)
	err = mysqlDB.Exec(sql).Error
	if err != nil {
		fmt.Println("0", "Exec failed ", err.Error(), sql)
		panic(err.Error())
	}
	sqlDB, _ := mysqlDB.DB()
	sqlDB.Close()

	dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		config.Config.Mysql.DBUserName, config.Config.Mysql.DBPassword, config.Config.Mysql.DBAddress[0], config.Config.Mysql.DBDatabaseName)

	mysqlDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("0", "Open failed ", err.Error(), dsn)
		panic(err.Error())
	}
	// Perform additional MySQL-specific setup or configurations
	sqlDB, _ = mysqlDB.DB()
	fmt.Println("open db ok ", dsn)

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
