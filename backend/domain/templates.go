package domain

import "errors"

var (
	ErrTemplateNotFound      = errors.New("template not found")
	ErrTemplateAlreadyExists = errors.New("template already exists")
)

type Replacement struct {
	From string
	To   string
}

type Template struct {
	ID           string
	Name         string
	Replacements []Replacement
}

type TemplateRepository interface {
	FindAll() ([]Template, error)
	FindByName(name string) (Template, error)
	FindByID(id string) (Template, error)
	StoreTemplate(template Template) error
	UpdateTemplate(id string, updateTemplate Template) error
	DeleteTemplate(id string) error
}
