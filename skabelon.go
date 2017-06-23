/*
Package skabelon implements total templating functionality
*/
package skabelon

import (
	"os"
	"strings"
	"text/template"
)

type Skabelon struct {
	values   map[string]string
	options  []string
	template *template.Template
}

func New() *Skabelon {
	instance := new(Skabelon)
	instance.options = []string{"missingkey=error"}
	instance.initSysEnv()
	return instance
}

func (instance *Skabelon) SetOptions(options []string) {
	instance.options = options
}

func (instance *Skabelon) initSysEnv() {
	instance.values = make(map[string]string)
	for _, tuple := range os.Environ() {
		pair := strings.Split(tuple, "=")
		instance.Set(pair[0], pair[1])
	}
}

func (instance *Skabelon) Set(key, value string) {
	instance.values[key] = value
}

func (instance *Skabelon) Get(key string) (string, bool) {
	val, exist := instance.values[key]
	return val, exist
}

func (instance *Skabelon) Parse(name, tmpl string) error {
	if instance.template == nil {
		instance.template = template.New(name).Option(instance.options...)
	}
	var tt *template.Template
	if instance.template.Name() == name {
		tt = instance.template
	} else {
		tt = instance.template.New(name)
	}
	_, err := tt.Parse(tmpl)
	return err
}

func (instance *Skabelon) Exec(file *os.File) error {
	return instance.template.Execute(file, instance.values)
}
