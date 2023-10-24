package database

import (
	_ "github.com/dialects/sqlite"
	"github.com/jinzhu/gorm"
)

var (
	DBConn *gorm.DB
)
