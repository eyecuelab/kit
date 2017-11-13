package psql

import "github.com/jinzhu/gorm"

type Scope func(db *gorm.DB) *gorm.DB
