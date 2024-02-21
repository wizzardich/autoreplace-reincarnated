package dto

import (
	"github.com/wizzardich/autoreplace-reincarnated/domain"
)

type Replacement struct {
	From string `bson:"from" json:"from"`
	To   string `bson:"to" json:"to"`
}

type Template struct {
	ID           string        `bson:"_id,omitempty" json:"id,omitempty"`
	Name         string        `bson:"name" json:"name"`
	Replacements []Replacement `bson:"reps" json:"replacements"`
}

type TemplateIdentifier struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func ReplacementFromDTO(dto Replacement) domain.Replacement {
	return domain.Replacement{
		From: dto.From,
		To:   dto.To,
	}
}

func TemplateFromDTO(dto Template) domain.Template {
	replacements := make([]domain.Replacement, len(dto.Replacements))
	for i, r := range dto.Replacements {
		replacements[i] = ReplacementFromDTO(r)
	}

	return domain.Template{
		ID:           dto.ID,
		Name:         dto.Name,
		Replacements: replacements,
	}
}

func ReplacementToDTO(replacement domain.Replacement) Replacement {
	return Replacement{
		From: replacement.From,
		To:   replacement.To,
	}
}

func TemplateToDTO(template domain.Template) Template {
	replacements := make([]Replacement, len(template.Replacements))
	for i, r := range template.Replacements {
		replacements[i] = ReplacementToDTO(r)
	}

	return Template{
		ID:           template.ID,
		Name:         template.Name,
		Replacements: replacements,
	}
}
