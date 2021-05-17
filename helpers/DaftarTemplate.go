package helpers

import (
	"errors"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

type DaftarTemplate struct {
	Templates map[string]*template.Template
}

func (t *DaftarTemplate) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := t.Templates[name]
	if !ok {
		err := errors.New("Template tidak ditemukan -->" + name)
		return err
	}
	return tmpl.ExecuteTemplate(w, "base.html", data)
}
