package api

import (
	"gorm.io/gorm"
)

type Container struct {
	DB *gorm.DB
}
