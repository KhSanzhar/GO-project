package handlers

import (
	"github.com/KhSanzar/Go-project/internal/models"
	"github.com/KhSanzar/Go-project/internal/render"
	"net/http"
)

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "home.page.gohtml", &models.TemplateData{})
}
