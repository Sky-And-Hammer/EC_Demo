package config

import (
	"html/template"
	"os"

	//  Golang Configuration tool that support YAML, JSON, Shell Environment
	"github.com/jinzhu/configor"
	//	a HTML sanitizer implemented in Go. It is fast and highly configurable.
	"github.com/microcosm-cc/bluemonday"

	"github.com/Sky-And-Hammer/render"
)

type SMTPConfig struct {
	HOST     string
	Port     string
	User     string
	Password string
	Site     string
}

var Config = struct {
	Port uint `default:"7000" env:"PORT"`
	DB   struct {
		Name     string `defualt:"ec_example"`
		Adapter  string `default:"mysql"`
		User     string
		Password string
	}
	SMTP SMTPConfig
}{}

var (
	Root = os.Getenv("GOPATH") + "/src/github.com/Sky-And-Hammer/EC_Demo"
	View *render.Render
)

func init() {
	if err := configor.Load(&Config, "config/database.yml", "config/smtp.yml"); err != nil {
		panic(err)
	}

	View = render.New()
	htmlSanitizer := bluemonday.UGCPolicy()
	View.RegisterFuncMap("raw", func(str string) template.HTML {
		return template.HTML(htmlSanitizer.Sanitize(str))
	})
}

func (s SMTPConfig) HostWithPort() string {
	return s.HOST + ":" + s.Port
}
