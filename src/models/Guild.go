package models

import (
	"errors"
	"strings"
	"time"
)

type Guild struct {
	ID          uint64    `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	ImageKey    string    `json:"imageKey,omitempty"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"createdAt,omitempty"`
	UpdatedAt   time.Time `json:"updatedAt,omitempty"`
}

func (g *Guild) Prepare() error {
	if prepareError := g.validate(); prepareError != nil {
		return prepareError
	}

	g.format()

	return nil
}

func (g *Guild) validate() error {
	if g.Name == "" {
		return errors.New("name is a mandatory parameter and should not be empty")
	}

	if g.ImageKey == "" {
		return errors.New("image Key is a mandatory parameter and should not be empty")
	}

	if g.Description == "" {
		return errors.New("description is a mandatory parameter and should not be empty")
	}

	return nil
}

func (g *Guild) format() {
	g.Name = strings.TrimSpace(g.Name)
	g.ImageKey = strings.TrimSpace(g.ImageKey)
	g.Description = strings.TrimSpace(g.Description)
}
