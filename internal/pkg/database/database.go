package database

import (
	"fmt"
	"github.com/haapjari/repository-database-api/internal/pkg/util"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Options struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
	TimeZone string
}

type Database struct {
	opt *Options
}

func NewDatabase(opt *Options) *Database {
	return &Database{
		opt: opt,
	}
}

func (d *Database) Connect() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		d.opt.Host, d.opt.Username, d.opt.Password, d.opt.Database, d.opt.Port, d.opt.TimeZone)), &gorm.Config{})
	if err != nil {
		return nil, util.Error(err)
	}

	return db, nil
}
