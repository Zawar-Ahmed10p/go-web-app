package config

import (
	"html/template"
	"log"
)

//AppConfig
type Appconfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
}
