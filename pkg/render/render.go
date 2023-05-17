package render

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parseTemplate, err := template.ParseFiles("./templates/" + tmpl)
	if err != nil {
		fmt.Println("Parsing error", err)
		return
	}
	err = parseTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Parsing error", err)
		return
	}
}
