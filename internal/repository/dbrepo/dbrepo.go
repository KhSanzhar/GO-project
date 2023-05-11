package dbrepo

import (
	"github.com/KhSanzar/Go-project/internal/config"
	"github.com/KhSanzar/Go-project/internal/repository"
	"gorm.io/gorm"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *gorm.DB
}

func NewPostgresRepo(conn *gorm.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: a,
		DB:  conn,
	}
}
