package controllers

import (
	. "mhilmi999/project-2-mhilmi999/middlewares"
	"net/http"

	"github.com/labstack/echo/v4"
)

// type M map[string]interface{}

// type Renderer struct{
// 	template 	*template.Template
// 	debug		bool
// 	location	string
// }

// func NewRenderer(location string, debug bool) *Renderer{
// 	tpl := new(Renderer)
// 	tpl.location = location
// 	tpl.debug = debug

// 	tpl.ReloadTemplates()

// 	return tpl
// }

// func (t *Renderer) ReloadTemplates() { // Memparsing template ketika inisialisasi objek renderer
// 	t.template = template.Must(template.ParseGlob(t.location))
// }

// func (t *Renderer) Render( // Me-render template yang telah diparsing sebagai output
// 	w io.Writer,
// 	name string,
// 	data interface{},
// 	c echo.Context,
// ) error {
// 	if t.debug{
// 		t.ReloadTemplates()
// 	}

// 	return t.template.ExecuteTemplate(w, name, data)
// }

func HomeView(c echo.Context) error {
	flash := c.(*CustomContext).GetFlash()
	return c.Render(http.StatusOK, "home.html", map[string]interface{}{
		"flash": flash,
	})
}
