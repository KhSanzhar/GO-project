package handlers

import (
	"github.com/KhSanzar/Go-project/internal/config"
	"github.com/KhSanzar/Go-project/internal/repository"
	"github.com/KhSanzar/Go-project/internal/repository/dbrepo"
	"gorm.io/gorm"
)

// Repo is repository used by handlers
var Repo *Repository

// Repository is repository type
type Repository struct {
	App *config.AppConfig
	DB  repository.DatabaseRepo
}

// NewRepo creates new repository
func NewRepo(a *config.AppConfig, db *gorm.DB) *Repository {
	return &Repository{
		App: a,
		DB:  dbrepo.NewPostgresRepo(db, a),
	}
}

// NewHandlers sets repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}
