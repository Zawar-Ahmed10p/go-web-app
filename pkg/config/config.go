package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

//AppConfig
type Appconfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
}
