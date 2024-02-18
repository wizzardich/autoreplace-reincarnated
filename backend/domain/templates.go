package domain

import "errors"

var ErrTemplateNotFound = errors.New("template not found")

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
}
