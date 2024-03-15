package domain

import "errors"

var (
	ErrTemplateNotFound      = errors.New("template not found")
	ErrTemplateAlreadyExists = errors.New("template already exists")
)

type IDString = string

type Replacement struct {
	From string
	To   string
}

type Template struct {
	ID           IDString
	Name         string
	Replacements []Replacement
}

type TemplateRepository interface {
	FindAll() ([]Template, error)
	FindByName(name string) (Template, error)
	FindByID(id IDString) (Template, error)
	StoreTemplate(template Template) (*IDString, error)
	UpdateTemplate(id IDString, updateTemplate Template) error
	DeleteTemplate(id IDString) error
}
