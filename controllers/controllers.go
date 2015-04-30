package controllers

import (
	"net/http"

	"github.com/unrolled/render"
)

var r *render.Render

func init() {
	r = render.New(render.Options{
		Directory:     "/Users/troy/Projects/txtrocket/templates",
		IsDevelopment: true,
		//      Layout: "layouts/layout",
		Extensions: []string{".html"},
		//      Funcs: []template.FuncMap{AssetHelpers(root)},
	})

}

// RenderHTML renders a template to ResponseWriter
func RenderHTML(w http.ResponseWriter, status int, name string, binding interface{}, layout string) {
	if layout != "" {
		r.HTML(w, status, name, binding, render.HTMLOptions{Layout: layout})
	} else {
		r.HTML(w, status, name, binding)
	}
}
