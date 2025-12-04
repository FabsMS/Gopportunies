package config

import "gorm.io/gorm"

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	return nil // nulo
}

func GetLogger(p string) {
	// Initialize Logger
	logger = NewLogger(p)
	return logger
}
